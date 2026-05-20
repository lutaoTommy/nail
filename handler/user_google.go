package handler

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"nail/config"
	"nail/language"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var googleOAuthHTTPClient = &http.Client{Timeout: 10 * time.Second}

type googleLoginBody struct {
	AccessToken string `json:"access_token"`
}

type googleTokenInfo struct {
	Aud              string `json:"aud"`
	Azp              string `json:"azp"`
	Sub              string `json:"sub"`
	Email            string `json:"email"`
	EmailVerified    string `json:"email_verified"`
	ExpiresIn        string `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type googleUserInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

func googleLoginHandler(ctx iris.Context) {
	var body googleLoginBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_INVALID_PARAM")})
		return
	}
	body.AccessToken = strings.TrimSpace(body.AccessToken)
	if body.AccessToken == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_GOOGLE_ACCESS_TOKEN")})
		return
	}
	ip := GetClientIP(ctx)
	sum := sha256.Sum256([]byte(body.AccessToken))
	accKey := "google:" + hex.EncodeToString(sum[:])
	if ok, retrySec := AllowLogin(ip, accKey); !ok {
		res := iris.Map{"result_code": 429, "result_msg": language.GetRawMessage("E_ACCOUNT_LOCKED")}
		if retrySec > 0 {
			res["retry_after"] = retrySec
		}
		ctx.JSON(res)
		return
	}
	token, err := googleLogin(ctx.Request().Context(), body.AccessToken)
	if err != nil {
		RecordLoginFailure(ip, accKey)
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	RecordLoginSuccess(ip, accKey)
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "token": token})
}

func googleLogin(ctx context.Context, rawAccessToken string) (sessionToken string, err error) {
	sub, emailNorm, vErr := validateGoogleAccessToken(ctx, rawAccessToken)
	if vErr != nil {
		return "", vErr
	}
	db := getMysqlConn()

	var bySub User
	err = db.Where("openid = ?", sub).First(&bySub).Error
	if err == nil {
		return finalizeOAuthSession(db, &bySub)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	var byEmail User
	err = db.Where("email = ?", emailNorm).First(&byEmail).Error
	if err == nil {
		if byEmail.OpenID != "" && byEmail.OpenID != sub {
			return "", newError(403, "E_GOOGLE_EMAIL_BOUND_OTHER")
		}
		if byEmail.Status != 1 {
			return "", newError(400, "E_INVALID_USER")
		}
		newTk, err := newToken()
		if err != nil {
			return "", err
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		res := db.Model(&User{}).Where("user_id = ?", byEmail.UserId).Updates(map[string]interface{}{
			"openid":     sub,
			"token":      newTk,
			"login_time": now,
		})
		if res.Error != nil {
			return "", res.Error
		}
		if res.RowsAffected == 0 {
			return "", newError(500, "E_INVALID_USER")
		}
		return newTk, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	// 自动开户
	pwdRand := make([]byte, 32)
	if _, err := rand.Read(pwdRand); err != nil {
		return "", err
	}
	hashed, err := HashPassword(base64.RawURLEncoding.EncodeToString(pwdRand))
	if err != nil {
		return "", err
	}
	tk, err := newToken()
	if err != nil {
		return "", err
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	u := User{
		OpenID:       sub,
		Email:        emailNorm,
		Passwd:       hashed,
		UserId:       RandStringBytes(5),
		Token:        tk,
		Status:       1,
		Nickname:     nicknameFromEmail(emailNorm),
		Language:     "en-US",
		RegisterTime: now,
		LoginTime:    now,
	}
	if err := db.Create(&u).Error; err != nil {
		var retry User
		if err2 := db.Where("openid = ?", sub).First(&retry).Error; err2 == nil {
			return finalizeOAuthSession(db, &retry)
		}
		return "", err
	}
	return tk, nil
}

func validateGoogleAccessToken(ctx context.Context, accessToken string) (sub, emailNorm string, err error) {
	audiences := config.GetGoogleOAuthClientIDs()
	if len(audiences) == 0 {
		return "", "", newError(503, "E_GOOGLE_NOT_CONFIGURED")
	}

	info, err := fetchGoogleTokenInfo(ctx, accessToken)
	if err != nil {
		return "", "", newError(401, "E_GOOGLE_ACCESS_TOKEN_INVALID")
	}
	if info.Error != "" || strings.TrimSpace(info.Sub) == "" {
		return "", "", newError(401, "E_GOOGLE_ACCESS_TOKEN_INVALID")
	}
	if exp := strings.TrimSpace(info.ExpiresIn); exp != "" {
		if sec, e := strconv.Atoi(exp); e == nil && sec <= 0 {
			return "", "", newError(401, "E_GOOGLE_ACCESS_TOKEN_INVALID")
		}
	}
	if !googleAudienceAllowed(info.Aud, info.Azp, audiences) {
		return "", "", newError(401, "E_GOOGLE_ACCESS_TOKEN_INVALID")
	}
	if !googleEmailVerifiedOK(info.EmailVerified) {
		return "", "", newError(401, "E_GOOGLE_EMAIL_NOT_VERIFIED")
	}

	sub = strings.TrimSpace(info.Sub)
	emailRaw := strings.TrimSpace(info.Email)
	if emailRaw == "" {
		ui, uErr := fetchGoogleUserInfo(ctx, accessToken)
		if uErr != nil || ui == nil {
			return "", "", newError(401, "E_GOOGLE_ACCESS_TOKEN_INVALID")
		}
		if strings.TrimSpace(ui.Sub) != "" && ui.Sub != sub {
			return "", "", newError(401, "E_GOOGLE_ACCESS_TOKEN_INVALID")
		}
		if !ui.EmailVerified {
			return "", "", newError(401, "E_GOOGLE_EMAIL_NOT_VERIFIED")
		}
		emailRaw = ui.Email
	}

	emailNorm, err = normalizeOAuthEmail(emailRaw)
	if err != nil {
		return "", "", err
	}
	if emailNorm == "" {
		return "", "", newError(401, "E_GOOGLE_EMAIL_NOT_VERIFIED")
	}
	return sub, emailNorm, nil
}

func fetchGoogleTokenInfo(ctx context.Context, accessToken string) (*googleTokenInfo, error) {
	u := "https://oauth2.googleapis.com/tokeninfo?access_token=" + url.QueryEscape(accessToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := googleOAuthHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tokeninfo status %d", resp.StatusCode)
	}
	var info googleTokenInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

func fetchGoogleUserInfo(ctx context.Context, accessToken string) (*googleUserInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := googleOAuthHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("userinfo status %d", resp.StatusCode)
	}
	var ui googleUserInfo
	if err := json.Unmarshal(body, &ui); err != nil {
		return nil, err
	}
	return &ui, nil
}

func googleAudienceAllowed(aud, azp string, allowed []string) bool {
	match := func(v string) bool {
		v = strings.TrimSpace(v)
		if v == "" {
			return false
		}
		for _, id := range allowed {
			if v == id {
				return true
			}
		}
		return false
	}
	return match(aud) || match(azp)
}

func googleEmailVerifiedOK(v string) bool {
	v = strings.TrimSpace(strings.ToLower(v))
	return v == "true" || v == "1"
}

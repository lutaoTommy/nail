package handler

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"time"

	"nail/config"
	"nail/language"

	"github.com/kataras/iris/v12"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

type googleLoginBody struct {
	IDToken string `json:"id_token"`
}

func googleLoginHandler(ctx iris.Context) {
	var body googleLoginBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_INVALID_PARAM")})
		return
	}
	body.IDToken = strings.TrimSpace(body.IDToken)
	if body.IDToken == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_GOOGLE_ID_TOKEN")})
		return
	}
	ip := GetClientIP(ctx)
	sum := sha256.Sum256([]byte(body.IDToken))
	accKey := "google:" + hex.EncodeToString(sum[:])
	if ok, retrySec := AllowLogin(ip, accKey); !ok {
		res := iris.Map{"result_code": 429, "result_msg": language.GetRawMessage("E_ACCOUNT_LOCKED")}
		if retrySec > 0 {
			res["retry_after"] = retrySec
		}
		ctx.JSON(res)
		return
	}
	token, err := googleLogin(ctx.Request().Context(), body.IDToken)
	if err != nil {
		RecordLoginFailure(ip, accKey)
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	RecordLoginSuccess(ip, accKey)
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "token": token})
}

func googleLogin(ctx context.Context, rawIDToken string) (sessionToken string, err error) {
	sub, emailNorm, vErr := validateGoogleIDToken(ctx, rawIDToken)
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

func validateGoogleIDToken(ctx context.Context, raw string) (sub, emailNorm string, err error) {
	audiences := config.GetGoogleOAuthClientIDs()
	if len(audiences) == 0 {
		return "", "", newError(503, "E_GOOGLE_NOT_CONFIGURED")
	}
	var payload *idtoken.Payload
	for _, aud := range audiences {
		p, e := idtoken.Validate(ctx, raw, aud)
		if e != nil {
			continue
		}
		payload = p
		break
	}
	if payload == nil {
		return "", "", newError(401, "E_GOOGLE_ID_TOKEN_INVALID")
	}
	sub = strings.TrimSpace(payload.Subject)
	if sub == "" {
		return "", "", newError(401, "E_GOOGLE_ID_TOKEN_INVALID")
	}
	emailRaw, _ := payload.Claims["email"].(string)
	emailNorm, err = normalizeOAuthEmail(emailRaw)
	if err != nil {
		return "", "", err
	}
	if emailNorm == "" {
		return "", "", newError(401, "E_GOOGLE_EMAIL_NOT_VERIFIED")
	}
	if !oauthClaimBool(payload.Claims, "email_verified") {
		return "", "", newError(401, "E_GOOGLE_EMAIL_NOT_VERIFIED")
	}
	return sub, emailNorm, nil
}

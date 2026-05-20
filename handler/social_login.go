package handler

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"sync"
	"time"

	"nail/config"
	"nail/language"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

const (
	socialChallengeTTL = 60 * time.Second
	socialReqMaxSkew   = 60 * time.Second
)

type socialChallengeReq struct {
	DeviceID string `json:"device_id"`
}

type socialChallengeRec struct {
	DeviceID  string
	IP        string
	ExpiresAt time.Time
	Used      bool
}

type socialLoginReq struct {
	Provider  string `json:"provider"`
	OpenID    string `json:"openid"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	Nonce     string `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	DeviceID  string `json:"device_id"`
	Signature string `json:"signature"`
}

var (
	socialChallengeMu sync.Mutex
	socialChallenges  = map[string]*socialChallengeRec{}
)

func socialChallengeHandler(ctx iris.Context) {
	var req socialChallengeReq
	_ = ctx.ReadJSON(&req)
	req.DeviceID = strings.TrimSpace(req.DeviceID)
	if req.DeviceID == "" {
		req.DeviceID = strings.TrimSpace(ctx.GetHeader("X-Device-ID"))
	}
	nonce, err := genSocialNonce()
	if err != nil {
		ctx.JSON(iris.Map{"result_code": 500, "result_msg": err.Error()})
		return
	}
	ip := GetClientIP(ctx)
	now := time.Now()
	exp := now.Add(socialChallengeTTL)
	socialChallengeMu.Lock()
	defer socialChallengeMu.Unlock()
	cleanupSocialChallengesLocked(now)
	socialChallenges[nonce] = &socialChallengeRec{
		DeviceID:  req.DeviceID,
		IP:        ip,
		ExpiresAt: exp,
		Used:      false,
	}
	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"nonce":       nonce,
		"expires_in":  int(socialChallengeTTL.Seconds()),
	})
}

func socialLoginHandler(ctx iris.Context) {
	secret := strings.TrimSpace(config.GetSocialAppSecret())
	if secret == "" {
		ctx.JSON(iris.Map{"result_code": 503, "result_msg": language.GetRawMessage("E_SOCIAL_NOT_CONFIGURED")})
		return
	}
	var req socialLoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_INVALID_PARAM")})
		return
	}
	req.Provider = strings.TrimSpace(strings.ToLower(req.Provider))
	req.OpenID = strings.TrimSpace(req.OpenID)
	req.Email = strings.TrimSpace(req.Email)
	req.Nickname = strings.TrimSpace(req.Nickname)
	req.Nonce = strings.TrimSpace(req.Nonce)
	req.DeviceID = strings.TrimSpace(req.DeviceID)
	req.Signature = strings.TrimSpace(req.Signature)
	if req.Provider == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_PROVIDER")})
		return
	}
	if req.Provider != "google" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_SOCIAL_PROVIDER_INVALID")})
		return
	}
	if req.OpenID == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_OPENID")})
		return
	}
	if req.Email == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_EMAIL")})
		return
	}
	if req.Nonce == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_NONCE")})
		return
	}
	if req.Timestamp == 0 {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_TIMESTAMP")})
		return
	}
	if req.Signature == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_SIGNATURE")})
		return
	}
	if req.DeviceID == "" {
		req.DeviceID = strings.TrimSpace(ctx.GetHeader("X-Device-ID"))
	}
	if req.DeviceID == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_NO_DEVICE_ID")})
		return
	}
	now := time.Now()
	ts := time.Unix(req.Timestamp, 0)
	if now.Sub(ts) > socialReqMaxSkew || ts.Sub(now) > socialReqMaxSkew {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": language.GetRawMessage("E_REQUEST_EXPIRED")})
		return
	}
	ip := GetClientIP(ctx)
	ch, err := consumeSocialChallenge(req.Nonce, req.DeviceID, ip, now)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	_ = ch
	payload := socialSignaturePayload(&req)
	if !verifySocialSignature(payload, req.Signature, secret) {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": language.GetRawMessage("E_SIGNATURE_INVALID")})
		return
	}

	token, err := socialLoginByGoogle(req.OpenID, req.Email, req.Nickname)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "token": token})
}

func socialLoginByGoogle(openid, email, nickname string) (string, error) {
	emailNorm, err := normalizeOAuthEmail(email)
	if err != nil {
		return "", err
	}
	if emailNorm == "" {
		return "", newError(400, "E_NO_EMAIL")
	}
	if nickname != "" && len([]rune(nickname)) > 100 {
		return "", newError(400, "E_TOO_LONG")
	}
	db := getMysqlConn()
	var bySub User
	err = db.Where("openid = ?", openid).First(&bySub).Error
	if err == nil {
		return finalizeOAuthSession(db, &bySub)
	}
	if err != gorm.ErrRecordNotFound {
		return "", err
	}
	var byEmail User
	err = db.Where("email = ?", emailNorm).First(&byEmail).Error
	if err == nil {
		if byEmail.OpenID != "" && byEmail.OpenID != openid {
			return "", newError(403, "E_SOCIAL_EMAIL_BOUND_OTHER")
		}
		if byEmail.Status != 1 {
			return "", newError(400, "E_INVALID_USER")
		}
		newTk, err := newToken()
		if err != nil {
			return "", err
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		data := map[string]interface{}{
			"openid":     openid,
			"token":      newTk,
			"login_time": now,
		}
		if byEmail.Nickname == "" && nickname != "" {
			data["nickname"] = nickname
		}
		res := db.Model(&User{}).Where("user_id = ?", byEmail.UserId).Updates(data)
		if res.Error != nil {
			return "", res.Error
		}
		if res.RowsAffected == 0 {
			return "", newError(500, "E_INVALID_USER")
		}
		return newTk, nil
	}
	if err != gorm.ErrRecordNotFound {
		return "", err
	}

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
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	finalNickname := nickname
	if finalNickname == "" {
		finalNickname = nicknameFromEmail(emailNorm)
	}
	u := User{
		OpenID:       openid,
		Email:        emailNorm,
		Passwd:       hashed,
		UserId:       RandStringBytes(5),
		Token:        tk,
		Status:       1,
		Nickname:     finalNickname,
		Language:     "en-US",
		RegisterTime: nowStr,
		LoginTime:    nowStr,
	}
	if err := db.Create(&u).Error; err != nil {
		var retry User
		if err2 := db.Where("openid = ?", openid).First(&retry).Error; err2 == nil {
			return finalizeOAuthSession(db, &retry)
		}
		return "", err
	}
	return tk, nil
}

func socialSignaturePayload(req *socialLoginReq) string {
	return "provider=" + req.Provider +
		"&openid=" + req.OpenID +
		"&email=" + strings.ToLower(req.Email) +
		"&nickname=" + req.Nickname +
		"&nonce=" + req.Nonce +
		"&timestamp=" + strconv.FormatInt(req.Timestamp, 10) +
		"&device_id=" + req.DeviceID
}

func verifySocialSignature(payload, gotSig, secret string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(payload))
	expected := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(gotSig))
}

func consumeSocialChallenge(nonce, deviceID, ip string, now time.Time) (*socialChallengeRec, error) {
	socialChallengeMu.Lock()
	defer socialChallengeMu.Unlock()
	cleanupSocialChallengesLocked(now)
	rec, ok := socialChallenges[nonce]
	if !ok {
		return nil, newError(401, "E_CHALLENGE_INVALID")
	}
	if now.After(rec.ExpiresAt) {
		delete(socialChallenges, nonce)
		return nil, newError(401, "E_CHALLENGE_EXPIRED")
	}
	if rec.Used {
		return nil, newError(401, "E_CHALLENGE_USED")
	}
	if rec.DeviceID != "" && rec.DeviceID != deviceID {
		return nil, newError(401, "E_CHALLENGE_INVALID")
	}
	if rec.IP != "" && rec.IP != ip {
		return nil, newError(401, "E_CHALLENGE_INVALID")
	}
	rec.Used = true
	return rec, nil
}

func cleanupSocialChallengesLocked(now time.Time) {
	for k, v := range socialChallenges {
		if v.Used || now.After(v.ExpiresAt) {
			delete(socialChallenges, k)
		}
	}
}

func genSocialNonce() (string, error) {
	b := make([]byte, 24)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

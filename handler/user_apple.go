package handler

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type appleLoginBody struct {
	IdentityToken string `json:"identity_token"`
}

func appleLoginHandler(ctx iris.Context) {
	var body appleLoginBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": Msg(ctx, "E_INVALID_PARAM")})
		return
	}
	body.IdentityToken = strings.TrimSpace(body.IdentityToken)
	if body.IdentityToken == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": Msg(ctx, "E_NO_APPLE_ID_TOKEN")})
		return
	}
	ip := GetClientIP(ctx)
	sum := sha256.Sum256([]byte(body.IdentityToken))
	accKey := "apple:" + hex.EncodeToString(sum[:])
	if ok, retrySec := AllowLogin(ip, accKey); !ok {
		ctx.JSON(loginLockedResponse(ctx, retrySec))
		return
	}
	token, err := appleLogin(ctx.Request().Context(), body.IdentityToken)
	if err != nil {
		ctx.JSON(loginFailureResponse(ctx, err, ip, accKey))
		return
	}
	RecordLoginSuccess(ip, accKey)
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "token": token})
}

func appleLogin(ctx context.Context, rawIdentityToken string) (sessionToken string, err error) {
	sub, emailNorm, vErr := validateAppleIdentityToken(ctx, rawIdentityToken)
	if vErr != nil {
		return "", vErr
	}
	db := getMysqlConn()

	var bySub User
	err = db.Where("apple_openid = ?", sub).First(&bySub).Error
	if err == nil {
		return finalizeOAuthSession(db, &bySub)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	if emailNorm != "" {
		var byEmail User
		err = db.Where("email = ?", emailNorm).First(&byEmail).Error
		if err == nil {
			if byEmail.AppleOpenID != "" && byEmail.AppleOpenID != sub {
				return "", newError(403, "E_APPLE_EMAIL_BOUND_OTHER")
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
				"apple_openid": sub,
				"token":        newTk,
				"login_time":   now,
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
	now := time.Now().Format("2006-01-02 15:04:05")
	u := User{
		AppleOpenID:  sub,
		Email:        emailNorm,
		Passwd:       hashed,
		UserId:       RandStringBytes(5),
		Token:        tk,
		Status:       1,
		Nickname:     "TINTA User",
		RegisterTime: now,
		LoginTime:    now,
	}
	if err := db.Create(&u).Error; err != nil {
		var retry User
		if err2 := db.Where("apple_openid = ?", sub).First(&retry).Error; err2 == nil {
			return finalizeOAuthSession(db, &retry)
		}
		return "", err
	}
	return tk, nil
}

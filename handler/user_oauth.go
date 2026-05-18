package handler

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

func finalizeOAuthSession(db *gorm.DB, u *User) (string, error) {
	if u.Status != 1 {
		return "", newError(400, "E_INVALID_USER")
	}
	newTk, err := newToken()
	if err != nil {
		return "", err
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	if err := db.Model(&User{}).Where("user_id = ?", u.UserId).Updates(map[string]interface{}{
		"token":      newTk,
		"login_time": now,
	}).Error; err != nil {
		return "", err
	}
	return newTk, nil
}

func nicknameFromEmail(email string) string {
	local := email
	if i := strings.IndexByte(email, '@'); i >= 0 {
		local = email[:i]
	}
	local = strings.TrimSpace(local)
	if local == "" {
		local = "user"
	}
	r := []rune(local)
	if len(r) > 100 {
		r = r[:100]
	}
	return string(r)
}

func oauthClaimBool(m map[string]interface{}, key string) bool {
	v, ok := m[key]
	if !ok {
		return false
	}
	switch t := v.(type) {
	case bool:
		return t
	case string:
		return strings.EqualFold(t, "true") || t == "1"
	case float64:
		return t != 0
	default:
		return false
	}
}

func normalizeOAuthEmail(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", nil
	}
	tmp := User{Email: strings.ToLower(raw)}
	if e := tmp.checkMail(); e != nil {
		return "", e
	}
	return strings.ToLower(raw), nil
}

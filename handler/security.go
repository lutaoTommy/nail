package handler

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// newToken 生成高强度 token（兼容 URL/HTTP Header，无 padding）。
// 32 bytes -> 43 chars (base64url, no padding)
func newToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// HashPassword 使用 bcrypt 对明文密码进行哈希存储。
// 注意：bcrypt 对输入长度超过 72 bytes 会截断，因此上层需要限制新密码长度。
func HashPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func isBcryptHash(s string) bool {
	// 常见前缀：$2a$ / $2b$ / $2y$
	return strings.HasPrefix(s, "$2a$") || strings.HasPrefix(s, "$2b$") || strings.HasPrefix(s, "$2y$")
}

// VerifyPassword 校验密码（仅支持 bcrypt 存储，不再兼容明文）。
func VerifyPassword(stored, plain string) (bool, error) {
	if stored == "" {
		return false, nil
	}
	// 若不是 bcrypt 格式，直接视为不匹配（脚本应已完成历史数据迁移）
	if !isBcryptHash(stored) {
		return false, nil
	}
	cerr := bcrypt.CompareHashAndPassword([]byte(stored), []byte(plain))
	if cerr == nil {
		return true, nil
	}
	// 密码不匹配是常见错误，视为校验失败而非系统错误
	if cerr == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	return false, cerr
}


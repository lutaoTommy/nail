package handler

import (
	"net"
	"strings"
	"sync"
	"time"

	"github.com/kataras/iris/v12"
)

const (
	loginWindow       = 15 * time.Minute
	loginBlock        = 15 * time.Minute
	loginMaxFailures  = 5
	vcodeWindow       = 15 * time.Minute
	vcodeBlock        = 15 * time.Minute
	vcodeMaxPerIP     = 10
	vcodeMaxPerEmail  = 5
)

type limitRec struct {
	Count        int
	WindowStart  time.Time
	BlockedUntil time.Time
	mu           sync.Mutex
}

var (
	loginLimitByIP     = make(map[string]*limitRec)
	loginLimitByAcc    = make(map[string]*limitRec)
	vcodeLimitByIP     = make(map[string]*limitRec)
	vcodeLimitByEmail  = make(map[string]*limitRec)
	loginMu            sync.RWMutex
	vcodeMu            sync.RWMutex
)

// GetClientIP 从请求中取客户端 IP（支持 X-Real-IP / X-Forwarded-For / RemoteAddr）
func GetClientIP(ctx iris.Context) string {
	if s := ctx.GetHeader("X-Real-IP"); s != "" {
		return strings.TrimSpace(strings.Split(s, ",")[0])
	}
	if s := ctx.GetHeader("X-Forwarded-For"); s != "" {
		return strings.TrimSpace(strings.Split(s, ",")[0])
	}
	addr := ctx.Request().RemoteAddr
	if host, _, err := net.SplitHostPort(addr); err == nil {
		return host
	}
	return addr
}

// AllowLogin 检查是否允许登录尝试（按 IP 与 账号 双重限制）
// accountKey 建议 "phone:xxx" 或 "email:xxx"
// 返回 allowed, retryAfterSeconds（被封时 >0）
func AllowLogin(ip, accountKey string) (allowed bool, retryAfterSec int) {
	now := time.Now()
	loginMu.Lock()
	defer loginMu.Unlock()

	check := func(m map[string]*limitRec, key string) (bool, int) {
		r, ok := m[key]
		if !ok {
			return true, 0
		}
		r.mu.Lock()
		defer r.mu.Unlock()
		if now.Before(r.BlockedUntil) {
			return false, int(time.Until(r.BlockedUntil).Seconds())
		}
		if now.Sub(r.WindowStart) > loginWindow {
			r.Count = 0
			r.WindowStart = now
		}
		return true, 0
	}

	ipAllowed, ipRetry := check(loginLimitByIP, ip)
	if !ipAllowed {
		return false, ipRetry
	}
	accAllowed, accRetry := check(loginLimitByAcc, accountKey)
	if !accAllowed {
		return false, accRetry
	}
	return true, 0
}

// RecordLoginFailure 记录一次登录失败（IP + 账号）
func RecordLoginFailure(ip, accountKey string) {
	now := time.Now()
	loginMu.Lock()
	defer loginMu.Unlock()

	record := func(m map[string]*limitRec, key string) {
		r, ok := m[key]
		if !ok {
			r = &limitRec{WindowStart: now}
			m[key] = r
		}
		r.mu.Lock()
		defer r.mu.Unlock()
		if now.Before(r.BlockedUntil) {
			return
		}
		if now.Sub(r.WindowStart) > loginWindow {
			r.Count = 0
			r.WindowStart = now
		}
		r.Count++
		if r.Count >= loginMaxFailures {
			r.BlockedUntil = now.Add(loginBlock)
		}
	}
	record(loginLimitByIP, ip)
	record(loginLimitByAcc, accountKey)
}

// RecordLoginSuccess 登录成功后清除该账号的失败计数（IP 计数保留，防止同 IP 多账号爆破）
func RecordLoginSuccess(ip, accountKey string) {
	loginMu.Lock()
	defer loginMu.Unlock()
	if r, ok := loginLimitByAcc[accountKey]; ok {
		r.mu.Lock()
		r.Count = 0
		r.mu.Unlock()
	}
}

// AllowVerificationRequest 检查是否允许请求邮箱验证码（按 IP + 邮箱 限制）
// 同一 IP 15 分钟内最多 10 次；同一邮箱 15 分钟内最多 5 次
func AllowVerificationRequest(ip, email string) (allowed bool, retryAfterSec int) {
	now := time.Now()
	vcodeMu.Lock()
	defer vcodeMu.Unlock()

	check := func(m map[string]*limitRec, key string, max int) (bool, int) {
		r, ok := m[key]
		if !ok {
			return true, 0
		}
		r.mu.Lock()
		defer r.mu.Unlock()
		if now.Before(r.BlockedUntil) {
			return false, int(time.Until(r.BlockedUntil).Seconds())
		}
		if now.Sub(r.WindowStart) > vcodeWindow {
			r.Count = 0
			r.WindowStart = now
		}
		if r.Count >= max {
			if r.BlockedUntil.Before(now) {
				r.BlockedUntil = now.Add(vcodeBlock)
			}
			return false, int(time.Until(r.BlockedUntil).Seconds())
		}
		return true, 0
	}

	ipAllowed, ipRetry := check(vcodeLimitByIP, ip, vcodeMaxPerIP)
	if !ipAllowed {
		return false, ipRetry
	}
	emailNorm := strings.ToLower(strings.TrimSpace(email))
	accAllowed, accRetry := check(vcodeLimitByEmail, emailNorm, vcodeMaxPerEmail)
	if !accAllowed {
		return false, accRetry
	}
	return true, 0
}

// RecordVerificationRequest 记录一次验证码请求（IP + 邮箱）
func RecordVerificationRequest(ip, email string) {
	now := time.Now()
	vcodeMu.Lock()
	defer vcodeMu.Unlock()

	record := func(m map[string]*limitRec, key string, max int) {
		r, ok := m[key]
		if !ok {
			r = &limitRec{WindowStart: now}
			m[key] = r
		}
		r.mu.Lock()
		defer r.mu.Unlock()
		if now.Before(r.BlockedUntil) {
			return
		}
		if now.Sub(r.WindowStart) > vcodeWindow {
			r.Count = 0
			r.WindowStart = now
		}
		r.Count++
		if r.Count >= max {
			r.BlockedUntil = now.Add(vcodeBlock)
		}
	}
	record(vcodeLimitByIP, ip, vcodeMaxPerIP)
	emailNorm := strings.ToLower(strings.TrimSpace(email))
	record(vcodeLimitByEmail, emailNorm, vcodeMaxPerEmail)
}

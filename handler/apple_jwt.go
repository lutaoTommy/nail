package handler

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"nail/config"

	"github.com/golang-jwt/jwt/v5"
)

const appleJWKSURL = "https://appleid.apple.com/auth/keys"
const appleIssuer = "https://appleid.apple.com"

type appleJWKS struct {
	Keys []appleJWK `json:"keys"`
}

type appleJWK struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

var (
	appleKeysMu      sync.RWMutex
	appleKeysCache   map[string]*rsa.PublicKey
	appleKeysExpires time.Time
)

func validateAppleIdentityToken(ctx context.Context, raw string) (sub, emailNorm string, err error) {
	sub, emailNorm, reason := verifyAppleIdentityToken(ctx, raw)
	if reason != "" {
		return "", "", newError(401, "E_APPLE_ID_TOKEN_INVALID")
	}
	return sub, emailNorm, nil
}

// DiagnoseAppleIdentityToken 返回验签失败原因（供日志排查，勿直接返回给客户端）。
func DiagnoseAppleIdentityToken(ctx context.Context, raw string) string {
	_, _, reason := verifyAppleIdentityToken(ctx, raw)
	if reason == "" {
		return "ok"
	}
	return reason
}

func verifyAppleIdentityToken(ctx context.Context, raw string) (sub, emailNorm, failReason string) {
	if !config.AppleSignInEnabled() {
		return "", "", "apple client_id/services_id not configured in config.ini"
	}
	allowedAud := config.GetAppleAllowedAudiences()
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", "", "empty identity token"
	}

	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithLeeway(60*time.Second),
	)
	token, err := parser.Parse(raw, func(t *jwt.Token) (interface{}, error) {
		kid, _ := t.Header["kid"].(string)
		if kid == "" {
			return nil, errors.New("missing kid")
		}
		return getApplePublicKey(ctx, kid)
	})
	if err != nil {
		hint := appleJWTUnverifiedHint(raw)
		return "", "", fmt.Sprintf("jwt parse: %v%s", err, hint)
	}
	if !token.Valid {
		return "", "", "jwt invalid"
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", "claims type invalid"
	}
	iss, _ := claims["iss"].(string)
	if iss != appleIssuer {
		return "", "", fmt.Sprintf("iss mismatch: got %q want %q", iss, appleIssuer)
	}

	audList := appleClaimAudiences(claims)
	audOK := false
	for _, aud := range audList {
		if appleAudienceAllowed(aud, allowedAud) {
			audOK = true
			break
		}
	}
	if !audOK {
		return "", "", fmt.Sprintf("aud %v not in allowed audiences client_id=%v services_id=%v",
			audList, config.GetAppleClientIDs(), config.GetAppleServicesIDs())
	}

	sub, _ = claims["sub"].(string)
	sub = strings.TrimSpace(sub)
	if sub == "" {
		return "", "", "missing sub"
	}

	emailRaw, _ := claims["email"].(string)
	emailNorm, err = normalizeOAuthEmail(emailRaw)
	if err != nil {
		return "", "", "email: " + err.Error()
	}
	if emailNorm != "" && !oauthClaimBool(claims, "email_verified") {
		return "", "", "email not verified"
	}
	return sub, emailNorm, ""
}

func appleClaimAudiences(claims jwt.MapClaims) []string {
	var out []string
	switch aud := claims["aud"].(type) {
	case string:
		if s := strings.TrimSpace(aud); s != "" {
			out = append(out, s)
		}
	case []interface{}:
		for _, a := range aud {
			if s, ok := a.(string); ok {
				s = strings.TrimSpace(s)
				if s != "" {
					out = append(out, s)
				}
			}
		}
	}
	return out
}

func appleJWTUnverifiedHint(raw string) string {
	parser := jwt.NewParser()
	tok, _, err := parser.ParseUnverified(raw, jwt.MapClaims{})
	if err != nil {
		return ""
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}
	iss, _ := claims["iss"].(string)
	aud := appleClaimAudiences(claims)
	exp, _ := claims["exp"]
	return fmt.Sprintf("; unverified iss=%q aud=%v exp=%v", iss, aud, exp)
}

func appleAudienceAllowed(aud string, allowed []string) bool {
	aud = strings.TrimSpace(aud)
	for _, id := range allowed {
		if aud == id {
			return true
		}
	}
	return false
}

func getApplePublicKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	appleKeysMu.RLock()
	if appleKeysCache != nil && time.Now().Before(appleKeysExpires) {
		if k, ok := appleKeysCache[kid]; ok {
			appleKeysMu.RUnlock()
			return k, nil
		}
	}
	appleKeysMu.RUnlock()

	if err := refreshAppleJWKS(ctx); err != nil {
		return nil, err
	}

	appleKeysMu.RLock()
	defer appleKeysMu.RUnlock()
	k, ok := appleKeysCache[kid]
	if !ok {
		return nil, fmt.Errorf("apple jwk kid not found: %s", kid)
	}
	return k, nil
}

func refreshAppleJWKS(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, appleJWKSURL, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("apple jwks http %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return err
	}
	var jwks appleJWKS
	if err := json.Unmarshal(body, &jwks); err != nil {
		return err
	}
	m := make(map[string]*rsa.PublicKey, len(jwks.Keys))
	for _, k := range jwks.Keys {
		if k.Kty != "RSA" || k.Kid == "" {
			continue
		}
		pub, err := rsaPublicKeyFromJWK(k.N, k.E)
		if err != nil {
			continue
		}
		m[k.Kid] = pub
	}
	if len(m) == 0 {
		return errors.New("apple jwks empty")
	}

	appleKeysMu.Lock()
	appleKeysCache = m
	appleKeysExpires = time.Now().Add(24 * time.Hour)
	appleKeysMu.Unlock()
	return nil
}

func rsaPublicKeyFromJWK(nB64, eB64 string) (*rsa.PublicKey, error) {
	nb, err := base64.RawURLEncoding.DecodeString(nB64)
	if err != nil {
		return nil, err
	}
	eb, err := base64.RawURLEncoding.DecodeString(eB64)
	if err != nil {
		return nil, err
	}
	var e int
	for _, b := range eb {
		e = e<<8 + int(b)
	}
	if e == 0 {
		e = 65537
	}
	n := new(big.Int).SetBytes(nb)
	return &rsa.PublicKey{N: n, E: e}, nil
}

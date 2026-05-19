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
	if !config.AppleSignInEnabled() {
		return "", "", newError(503, "E_APPLE_NOT_CONFIGURED")
	}
	clientIDs := config.GetAppleClientIDs()
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", "", newError(401, "E_APPLE_ID_TOKEN_INVALID")
	}

	parser := jwt.NewParser(jwt.WithValidMethods([]string{"RS256"}))
	token, err := parser.Parse(raw, func(t *jwt.Token) (interface{}, error) {
		kid, _ := t.Header["kid"].(string)
		if kid == "" {
			return nil, errors.New("missing kid")
		}
		return getApplePublicKey(ctx, kid)
	})
	if err != nil || !token.Valid {
		return "", "", newError(401, "E_APPLE_ID_TOKEN_INVALID")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", newError(401, "E_APPLE_ID_TOKEN_INVALID")
	}
	if iss, _ := claims["iss"].(string); iss != appleIssuer {
		return "", "", newError(401, "E_APPLE_ID_TOKEN_INVALID")
	}
	audOK := false
	switch aud := claims["aud"].(type) {
	case string:
		audOK = appleAudienceAllowed(aud, clientIDs)
	case []interface{}:
		for _, a := range aud {
			if s, ok := a.(string); ok && appleAudienceAllowed(s, clientIDs) {
				audOK = true
				break
			}
		}
	}
	if !audOK {
		return "", "", newError(401, "E_APPLE_ID_TOKEN_INVALID")
	}

	sub, _ = claims["sub"].(string)
	sub = strings.TrimSpace(sub)
	if sub == "" {
		return "", "", newError(401, "E_APPLE_ID_TOKEN_INVALID")
	}

	emailRaw, _ := claims["email"].(string)
	emailNorm, err = normalizeOAuthEmail(emailRaw)
	if err != nil {
		return "", "", err
	}
	if emailNorm != "" && !oauthClaimBool(claims, "email_verified") {
		return "", "", newError(401, "E_APPLE_EMAIL_NOT_VERIFIED")
	}
	return sub, emailNorm, nil
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

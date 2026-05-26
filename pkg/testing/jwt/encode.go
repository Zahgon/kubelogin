package jwt

import (
	"crypto/rsa"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

var PrivateKey = generateKey(1024)

func generateKey(b int) *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

type Claims struct {
	jwt.RegisteredClaims
	// aud claim is either a string or an array of strings.
	// https://tools.ietf.org/html/rfc7519#section-4.1.3
	Audience      []string `json:"aud,omitempty"`
	Nonce         string   `json:"nonce,omitempty"`
	Groups        []string `json:"groups,omitempty"`
	EmailVerified bool     `json:"email_verified,omitempty"`
}

func Encode(t *testing.T, claims Claims) string { _ = "STUB: not implemented"; return "" }

func EncodeF(t *testing.T, mutation func(*Claims)) string { _ = "STUB: not implemented"; return "" }

package oidc

import (
	"github.com/int128/kubelogin/pkg/jwt"
)

// Provider represents an OIDC provider.
type Provider struct {
	IssuerURL      string
	ClientID       string
	ClientSecret   string   // optional
	ExtraScopes    []string // optional
	RedirectURL    string   // optional
	PKCEMethod     PKCEMethod
	UseAccessToken bool
	RequestHeaders map[string]string
}

// PKCEMethod represents a preferred method of PKCE.
type PKCEMethod int

const (
	PKCEMethodAuto PKCEMethod = iota
	PKCEMethodNo
	PKCEMethodS256
)

// TokenSet represents a set of ID token and refresh token.
type TokenSet struct {
	IDToken      string
	RefreshToken string
}

func (ts TokenSet) DecodeWithoutVerify() (*jwt.Claims, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewState() (string, error) { _ = "STUB: not implemented"; return "", nil }

func NewNonce() (string, error) { _ = "STUB: not implemented"; return "", nil }

func random32() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func base64URLEncode(b []byte) string { _ = "STUB: not implemented"; return "" }

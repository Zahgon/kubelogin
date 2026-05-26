package client

import (
	"context"

	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/pkce"
	"golang.org/x/oauth2"
)

type AuthCodeURLInput struct {
	State                  string
	Nonce                  string
	PKCEParams             pkce.Params
	AuthRequestExtraParams map[string]string
}

type ExchangeAuthCodeInput struct {
	Code       string
	PKCEParams pkce.Params
	Nonce      string
}

type GetTokenByAuthCodeInput struct {
	BindAddress            []string
	State                  string
	Nonce                  string
	PKCEParams             pkce.Params
	AuthRequestExtraParams map[string]string
	LocalServerSuccessHTML string
	LocalServerCertFile    string
	LocalServerKeyFile     string
}

func (c *client) NegotiatedPKCEMethod() pkce.Method {
	_ = "STUB: not implemented"
	return *new(pkce.Method)
}

// GetTokenByAuthCode performs the authorization code flow.
func (c *client) GetTokenByAuthCode(ctx context.Context, in GetTokenByAuthCodeInput, localServerReadyChan chan<- string) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAuthCodeURL returns the URL of authentication request for the authorization code flow.
func (c *client) GetAuthCodeURL(in AuthCodeURLInput) string { _ = "STUB: not implemented"; return "" }

// ExchangeAuthCode exchanges the authorization code and token.
func (c *client) ExchangeAuthCode(ctx context.Context, in ExchangeAuthCodeInput) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authorizationRequestOptions(nonce string, pkceParams pkce.Params, extraParams map[string]string) []oauth2.AuthCodeOption {
	_ = "STUB: not implemented"
	return nil
}

func tokenRequestOptions(pkceParams pkce.Params) []oauth2.AuthCodeOption {
	_ = "STUB: not implemented"
	return nil
}

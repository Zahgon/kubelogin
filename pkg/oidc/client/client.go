package client

import (
	"context"
	"net/http"

	gooidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/int128/kubelogin/pkg/infrastructure/clock"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/pkce"
	"github.com/int128/oauth2dev"
	"golang.org/x/oauth2"
)

type Interface interface {
	GetAuthCodeURL(in AuthCodeURLInput) string
	ExchangeAuthCode(ctx context.Context, in ExchangeAuthCodeInput) (*oidc.TokenSet, error)
	GetTokenByAuthCode(ctx context.Context, in GetTokenByAuthCodeInput, localServerReadyChan chan<- string) (*oidc.TokenSet, error)
	NegotiatedPKCEMethod() pkce.Method
	GetTokenByROPC(ctx context.Context, username, password string) (*oidc.TokenSet, error)
	GetTokenByClientCredentials(ctx context.Context, in GetTokenByClientCredentialsInput) (*oidc.TokenSet, error)
	GetDeviceAuthorization(ctx context.Context) (*oauth2dev.AuthorizationResponse, error)
	ExchangeDeviceCode(ctx context.Context, authResponse *oauth2dev.AuthorizationResponse) (*oidc.TokenSet, error)
	Refresh(ctx context.Context, refreshToken string) (*oidc.TokenSet, error)
}

type client struct {
	httpClient           *http.Client
	provider             *gooidc.Provider
	oauth2Config         oauth2.Config
	clock                clock.Interface
	logger               logger.Interface
	negotiatedPKCEMethod pkce.Method
	useAccessToken       bool
}

func (c *client) wrapContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Refresh sends a refresh token request and returns a token set.
func (c *client) Refresh(ctx context.Context, refreshToken string) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verifyToken verifies the token with the certificates of the provider and the nonce.
// If the nonce is an empty string, it does not verify the nonce.
func (c *client) verifyToken(ctx context.Context, token *oauth2.Token, nonce string) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We intentionally do not perform a ClientID check here because there
// are some use cases in access_tokens where we *expect* the audience
// to differ. For example, one can explicitly set
// `audience=CLUSTER_CLIENT_ID` as an extra auth parameter.

// There is no `nonce` to check on the `access_token`. We rely on the
// above `nonce` check on the `id_token`.

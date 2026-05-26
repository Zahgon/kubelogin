package client

import (
	"context"

	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/oauth2dev"
)

// GetDeviceAuthorization initializes the device authorization code challenge
func (c *client) GetDeviceAuthorization(ctx context.Context) (*oauth2dev.AuthorizationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExchangeDeviceCode exchanges the device authorization code for an oidc.TokenSet
func (c *client) ExchangeDeviceCode(ctx context.Context, authResponse *oauth2dev.AuthorizationResponse) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

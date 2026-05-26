package client

import (
	"context"

	"github.com/int128/kubelogin/pkg/oidc"
)

type GetTokenByClientCredentialsInput struct {
	EndpointParams map[string][]string
}

// GetTokenByClientCredentials performs the client credentials flow.
func (c *client) GetTokenByClientCredentials(ctx context.Context, in GetTokenByClientCredentialsInput) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package client

import (
	"context"

	"github.com/int128/kubelogin/pkg/oidc"
)

// GetTokenByROPC performs the resource owner password credentials flow.
func (c *client) GetTokenByROPC(ctx context.Context, username, password string) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package clientcredentials

import (
	"context"

	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/oidc/client"
)

// DeviceCode provides the oauth2 device code flow.
type ClientCredentials struct {
	Logger logger.Interface
}

func (u *ClientCredentials) Do(ctx context.Context, in *client.GetTokenByClientCredentialsInput, oidcClient client.Interface) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

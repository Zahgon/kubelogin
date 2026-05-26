package ropc

import (
	"context"

	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/infrastructure/reader"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/oidc/client"
)

const usernamePrompt = "Username: "
const passwordPrompt = "Password: "

type Option struct {
	Username string
	Password string // If empty, read a password using Reader.ReadPassword()
}

// ROPC provides the resource owner password credentials flow.
type ROPC struct {
	Reader reader.Interface
	Logger logger.Interface
}

func (u *ROPC) Do(ctx context.Context, in *Option, oidcClient client.Interface) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

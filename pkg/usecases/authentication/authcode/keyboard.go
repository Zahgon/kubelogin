package authcode

import (
	"context"

	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/infrastructure/reader"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/oidc/client"
)

const keyboardPrompt = "Enter code: "

type KeyboardOption struct {
	AuthRequestExtraParams map[string]string
}

// Keyboard provides the authorization code flow with keyboard interactive.
type Keyboard struct {
	Reader reader.Interface
	Logger logger.Interface
}

func (u *Keyboard) Do(ctx context.Context, o *KeyboardOption, oidcClient client.Interface) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

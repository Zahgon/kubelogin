package devicecode

import (
	"context"

	"github.com/int128/kubelogin/pkg/infrastructure/browser"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/oidc/client"
)

type Option struct {
	SkipOpenBrowser bool
	BrowserCommand  string
}

// DeviceCode provides the oauth2 device code flow.
type DeviceCode struct {
	Browser browser.Interface
	Logger  logger.Interface
}

func (u *DeviceCode) Do(ctx context.Context, in *Option, oidcClient client.Interface) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *DeviceCode) openURL(ctx context.Context, o *Option, url string) {
	_ = "STUB: not implemented"
	return
}

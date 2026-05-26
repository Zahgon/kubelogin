package authcode

import (
	"context"
	"time"

	"github.com/int128/kubelogin/pkg/infrastructure/browser"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/oidc/client"
)

type BrowserOption struct {
	SkipOpenBrowser            bool
	BrowserCommand             string
	BindAddress                []string
	AuthenticationTimeout      time.Duration
	OpenURLAfterAuthentication string
	AuthRequestExtraParams     map[string]string
	LocalServerCertFile        string
	LocalServerKeyFile         string
}

// Browser provides the authentication code flow using the browser.
type Browser struct {
	Browser browser.Interface
	Logger  logger.Interface
}

func (u *Browser) Do(ctx context.Context, o *BrowserOption, oidcClient client.Interface) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *Browser) openURL(ctx context.Context, o *BrowserOption, url string) {
	_ = "STUB: not implemented"
	return
}

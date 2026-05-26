package cmd

import (
	"strings"

	"github.com/spf13/pflag"

	"github.com/int128/kubelogin/pkg/usecases/authentication"
)

type authenticationOptions struct {
	GrantType                  string
	ListenAddress              []string
	AuthenticationTimeoutSec   int
	SkipOpenBrowser            bool
	BrowserCommand             string
	LocalServerCertFile        string
	LocalServerKeyFile         string
	OpenURLAfterAuthentication string
	AuthRequestExtraParams     map[string]string
	Username                   string
	Password                   string
}

var allGrantType = strings.Join([]string{
	"auto",
	"authcode",
	"authcode-keyboard",
	"password",
	"device-code",
	"client-credentials",
}, "|")

func (o *authenticationOptions) addFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func (o *authenticationOptions) expandHomedir() { _ = "STUB: not implemented"; return }

func (o *authenticationOptions) grantOptionSet() (s authentication.GrantOptionSet, err error) {
	_ = "STUB: not implemented"
	return *new(authentication.GrantOptionSet), nil
}

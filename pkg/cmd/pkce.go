package cmd

import (
	"strings"

	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/spf13/pflag"
)

var allPKCEMethods = strings.Join([]string{"auto", "no", "S256"}, "|")

type pkceOptions struct {
	UsePKCE    bool
	PKCEMethod string
}

func (o *pkceOptions) addFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func (o *pkceOptions) pkceMethod() (oidc.PKCEMethod, error) {
	_ = "STUB: not implemented"
	return *new(oidc.PKCEMethod), nil
}

package cmd

import (
	"strings"

	"github.com/int128/kubelogin/pkg/tokencache"
	"github.com/spf13/pflag"
)

func getDefaultTokenCacheDir() string {
	_ = "STUB: not implemented"
	// https://github.com/int128/kubelogin/pull/975
	return ""
}

var allTokenCacheStorage = strings.Join([]string{"disk", "keyring", "none"}, "|")

type tokenCacheOptions struct {
	TokenCacheDir     string
	TokenCacheStorage string
}

func (o *tokenCacheOptions) addFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func (o *tokenCacheOptions) expandHomedir() { _ = "STUB: not implemented"; return }

func (o *tokenCacheOptions) tokenCacheConfig() (tokencache.Config, error) {
	_ = "STUB: not implemented"
	return *new(tokencache.Config), nil
}

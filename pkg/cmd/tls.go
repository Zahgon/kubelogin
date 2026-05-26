package cmd

import (
	"crypto/tls"

	"github.com/int128/kubelogin/pkg/tlsclientconfig"
	"github.com/spf13/pflag"
)

type tlsOptions struct {
	CACertFilename            []string
	CACertData                []string
	SkipTLSVerify             bool
	RenegotiateOnceAsClient   bool
	RenegotiateFreelyAsClient bool
}

func (o *tlsOptions) addFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func (o *tlsOptions) expandHomedir() { _ = "STUB: not implemented"; return }

func (o tlsOptions) tlsClientConfig() tlsclientconfig.Config {
	_ = "STUB: not implemented"
	return *new(tlsclientconfig.Config)
}

func (o tlsOptions) renegotiationSupport() tls.RenegotiationSupport {
	_ = "STUB: not implemented"
	return *new(tls.RenegotiationSupport)
}

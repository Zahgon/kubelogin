package cmd

import (
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/usecases/standalone"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const rootDescription = `Log in to the OpenID Connect provider.

You need to set up the OIDC provider, role binding, Kubernetes API server and kubeconfig.
To show the setup instruction:

	kubectl oidc-login setup

See https://github.com/int128/kubelogin for more.
`

// rootOptions represents the options for the root command.
type rootOptions struct {
	Kubeconfig            string
	Context               string
	User                  string
	tlsOptions            tlsOptions
	authenticationOptions authenticationOptions
}

func (o *rootOptions) addFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Root struct {
	Standalone standalone.Interface
	Logger     logger.Interface
}

func (cmd *Root) New() *cobra.Command { _ = "STUB: not implemented"; return nil }

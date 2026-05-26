package cmd

import (
	_ "embed"

	"github.com/int128/kubelogin/pkg/usecases/setup"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// setupOptions represents the options for setup command.
type setupOptions struct {
	IssuerURL             string
	ClientID              string
	ClientSecret          string
	RedirectURL           string
	ExtraScopes           []string
	UseAccessToken        bool
	RequestHeaders        map[string]string
	tlsOptions            tlsOptions
	pkceOptions           pkceOptions
	authenticationOptions authenticationOptions
}

func (o *setupOptions) addFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Setup struct {
	Setup setup.Interface
}

//go:embed setup.md
var setupLongDescription string

func (cmd *Setup) New() *cobra.Command { _ = "STUB: not implemented"; return nil }

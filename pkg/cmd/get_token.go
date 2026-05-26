package cmd

import (
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/usecases/credentialplugin"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// getTokenOptions represents the options for get-token command.
type getTokenOptions struct {
	IssuerURL             string
	ClientID              string
	ClientSecret          string
	RedirectURL           string
	ExtraScopes           []string
	UseAccessToken        bool
	RequestHeaders        map[string]string
	tokenCacheOptions     tokenCacheOptions
	tlsOptions            tlsOptions
	pkceOptions           pkceOptions
	authenticationOptions authenticationOptions
	ForceRefresh          bool
}

func (o *getTokenOptions) addFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func (o *getTokenOptions) expandHomedir() { _ = "STUB: not implemented"; return }

type GetToken struct {
	GetToken credentialplugin.Interface
	Logger   logger.Interface
}

func (cmd *GetToken) New() *cobra.Command { _ = "STUB: not implemented"; return nil }

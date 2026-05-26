package cmd

import (
	"github.com/int128/kubelogin/pkg/usecases/clean"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type cleanOptions struct {
	TokenCacheDir string
}

func (o *cleanOptions) addFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Clean struct {
	Clean clean.Interface
}

func (cmd *Clean) New() *cobra.Command { _ = "STUB: not implemented"; return nil }

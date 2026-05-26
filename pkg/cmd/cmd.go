package cmd

import (
	"context"

	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
)

// Set provides an implementation and interface for Cmd.
var Set = wire.NewSet(
	wire.Struct(new(Cmd), "*"),
	wire.Bind(new(Interface), new(*Cmd)),
	wire.Struct(new(Root), "*"),
	wire.Struct(new(GetToken), "*"),
	wire.Struct(new(Setup), "*"),
	wire.Struct(new(Clean), "*"),
)

type Interface interface {
	Run(ctx context.Context, args []string, version string) int
}

var defaultListenAddress = []string{"127.0.0.1:8000", "127.0.0.1:18000"}

const defaultAuthenticationTimeoutSec = 180

// Cmd provides interaction with command line interface (CLI).
type Cmd struct {
	Root     *Root
	GetToken *GetToken
	Setup    *Setup
	Clean    *Clean
	Logger   logger.Interface
}

// Run parses the command line arguments and executes the specified use-case.
// It returns an exit code, that is 0 on success or 1 on error.
func (cmd *Cmd) Run(ctx context.Context, args []string, version string) int {
	_ = "STUB: not implemented"
	return 0
}

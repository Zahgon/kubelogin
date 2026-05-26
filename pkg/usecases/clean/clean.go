package clean

import (
	"context"

	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/tokencache/repository"
)

var Set = wire.NewSet(
	wire.Struct(new(Clean), "*"),
	wire.Bind(new(Interface), new(*Clean)),
)

type Interface interface {
	Do(ctx context.Context, in Input) error
}

// Input represents an input of the Clean use-case.
type Input struct {
	TokenCacheDir string
}

type Clean struct {
	TokenCacheRepository repository.Interface
	Logger               logger.Interface
}

func (u *Clean) Do(ctx context.Context, in Input) error { _ = "STUB: not implemented"; return nil }

// Do not return an error because the keyring may not be available.

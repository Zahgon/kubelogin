// Package client provides a client of OpenID Connect.
package client

import (
	"context"

	gooidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/infrastructure/clock"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/pkce"
	"github.com/int128/kubelogin/pkg/tlsclientconfig"
	"github.com/int128/kubelogin/pkg/tlsclientconfig/loader"
)

var Set = wire.NewSet(
	wire.Struct(new(Factory), "*"),
	wire.Bind(new(FactoryInterface), new(*Factory)),
)

type FactoryInterface interface {
	New(ctx context.Context, prov oidc.Provider, tlsClientConfig tlsclientconfig.Config) (Interface, error)
}

type Factory struct {
	Loader loader.Loader
	Clock  clock.Interface
	Logger logger.Interface
}

// New returns an instance of infrastructure.Interface with the given configuration.
func (f *Factory) New(ctx context.Context, prov oidc.Provider, tlsClientConfig tlsclientconfig.Config) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func determinePKCEMethod(supportedMethods []string, preferredMethod oidc.PKCEMethod) pkce.Method {
	_ = "STUB: not implemented"
	return *new(pkce.Method)
}

func extractSupportedPKCEMethods(provider *gooidc.Provider) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

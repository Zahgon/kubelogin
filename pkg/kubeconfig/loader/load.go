package loader

import (
	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/kubeconfig"
	"k8s.io/client-go/tools/clientcmd/api"
)

var Set = wire.NewSet(
	wire.Struct(new(Loader), "*"),
	wire.Bind(new(Interface), new(*Loader)),
)

type Interface interface {
	GetCurrentAuthProvider(explicitFilename string, contextName kubeconfig.ContextName, userName kubeconfig.UserName) (*kubeconfig.AuthProvider, error)
}

type Loader struct{}

func (Loader) GetCurrentAuthProvider(explicitFilename string, contextName kubeconfig.ContextName, userName kubeconfig.UserName) (*kubeconfig.AuthProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadByDefaultRules(explicitFilename string) (*api.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findCurrentAuthProvider resolves the current auth provider.
// If contextName is given, this returns the user of the context.
// If userName is given, this ignores the context and returns the user.
// If any context or user is not found, this returns an error.
func findCurrentAuthProvider(config *api.Config, contextName kubeconfig.ContextName, userName kubeconfig.UserName) (*kubeconfig.AuthProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

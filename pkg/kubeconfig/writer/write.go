package writer

import (
	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/kubeconfig"
)

var Set = wire.NewSet(
	wire.Struct(new(Writer), "*"),
	wire.Bind(new(Interface), new(*Writer)),
)

type Interface interface {
	UpdateAuthProvider(p kubeconfig.AuthProvider) error
}

type Writer struct{}

func (Writer) UpdateAuthProvider(p kubeconfig.AuthProvider) error {
	_ = "STUB: not implemented"
	return nil
}

func copyAuthProviderConfig(p kubeconfig.AuthProvider, m map[string]string) {
	_ = "STUB: not implemented"
	return
}

func setOrDeleteKey(m map[string]string, key, value string) { _ = "STUB: not implemented"; return }

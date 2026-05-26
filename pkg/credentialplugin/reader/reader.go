// Package reader provides a loader for the credential plugin.
package reader

import (
	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/credentialplugin"
)

var Set = wire.NewSet(
	wire.Struct(new(Reader), "*"),
	wire.Bind(new(Interface), new(*Reader)),
)

type Interface interface {
	Read() (credentialplugin.Input, error)
}

type Reader struct{}

// Read parses the environment variable KUBERNETES_EXEC_INFO.
// If the environment variable is not given by kubectl, Read returns a zero value.
func (r Reader) Read() (credentialplugin.Input, error) {
	_ = "STUB: not implemented"
	return *new(credentialplugin.Input), nil
}

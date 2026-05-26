// Package writer provides a writer for the credential plugin.
package writer

import (
	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/credentialplugin"
	"github.com/int128/kubelogin/pkg/infrastructure/stdio"
)

var Set = wire.NewSet(
	wire.Struct(new(Writer), "*"),
	wire.Bind(new(Interface), new(*Writer)),
)

type Interface interface {
	Write(out credentialplugin.Output) error
}

type Writer struct {
	Stdout stdio.Stdout
}

// Write writes the ExecCredential to standard output for kubectl.
func (w *Writer) Write(out credentialplugin.Output) error { _ = "STUB: not implemented"; return nil }

func generateExecCredential(out credentialplugin.Output) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// If the API version is not available, fall back to v1beta1.

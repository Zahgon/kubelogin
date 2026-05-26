// Package reader provides the reader of standard input.
package reader

import (
	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/infrastructure/stdio"
)

// Set provides an implementation and interface for Reader.
var Set = wire.NewSet(
	wire.Struct(new(Reader), "*"),
	wire.Bind(new(Interface), new(*Reader)),
)

type Interface interface {
	ReadString(prompt string) (string, error)
	ReadPassword(prompt string) (string, error)
}

type Reader struct {
	Stdin stdio.Stdin
}

// ReadString reads a string from the stdin.
func (x *Reader) ReadString(prompt string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ReadPassword reads a password from the stdin without echo back.
func (*Reader) ReadPassword(prompt string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

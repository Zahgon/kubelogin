// Package loader provides loading certificates from files or base64 encoded string.
package loader

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/tlsclientconfig"
)

// Set provides an implementation and interface.
var Set = wire.NewSet(
	wire.Struct(new(Loader), "*"),
	wire.Bind(new(Interface), new(*Loader)),
)

type Interface interface {
	Load(config tlsclientconfig.Config) (*tls.Config, error)
}

// Loader represents a pool of certificates.
type Loader struct{}

func (l *Loader) Load(config tlsclientconfig.Config) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if empty, use the host's root CA set

func addFile(p *x509.CertPool, filename string) error { _ = "STUB: not implemented"; return nil }

func addBase64Encoded(p *x509.CertPool, s string) error { _ = "STUB: not implemented"; return nil }

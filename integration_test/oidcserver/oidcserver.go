// Package oidcserver provides a stub of OpenID Connect provider.
package oidcserver

import (
	"net/http"
	"testing"

	"github.com/int128/kubelogin/integration_test/keypair"
	"github.com/int128/kubelogin/integration_test/oidcserver/service"
	"github.com/int128/kubelogin/integration_test/oidcserver/testconfig"
)

// New starts a server for the OpenID Connect provider.
func New(t *testing.T, kp keypair.KeyPair, config testconfig.Config) service.Service {
	_ = "STUB: not implemented"
	return *new(service.Service)
}

func startServer(t *testing.T, h http.Handler, kp keypair.KeyPair) string {
	_ = "STUB: not implemented"
	return ""
}

// Unfortunately, httptest package did not work with keypair.KeyPair.
// We use httptest package only for allocating a new port.

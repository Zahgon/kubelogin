// Package handler provides HTTP handlers for the OpenID Connect Provider.
package handler

import (
	"net/http"
	"testing"

	"github.com/int128/kubelogin/integration_test/oidcserver/service"
)

func Register(t *testing.T, mux *http.ServeMux, provider service.Provider) {
	_ = "STUB: not implemented"
	return
}

// Handlers provides HTTP handlers for the OpenID Connect Provider.
// You need to implement the Provider interface.
// Note that this skips some security checks and is only for testing.
type Handlers struct {
	t        *testing.T
	provider service.Provider
}

func (h *Handlers) handleError(w http.ResponseWriter, r *http.Request, f func() error) {
	_ = "STUB: not implemented"
	return
}

func (h *Handlers) Discovery(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *Handlers) GetCertificates(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *Handlers) AuthenticateCode(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *Handlers) Exchange(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// 4.3. Resource Owner Password Credentials Grant
// https://tools.ietf.org/html/rfc6749#section-4.3

// 12.1. Refresh Request
// https://openid.net/specs/openid-connect-core-1_0.html#RefreshingAccessToken

// 5.2. Error Response
// https://tools.ietf.org/html/rfc6749#section-5.2

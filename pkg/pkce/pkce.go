// Package pkce provides generation of the PKCE parameters.
// See also https://tools.ietf.org/html/rfc7636.
package pkce

import "golang.org/x/oauth2"

type Method int

const (
	// Code challenge methods defined as https://tools.ietf.org/html/rfc7636#section-4.3
	NoMethod Method = iota
	MethodS256
)

// Params represents a set of the PKCE parameters.
type Params struct {
	Method   Method
	Verifier string
}

func (params Params) AuthCodeOption() oauth2.AuthCodeOption {
	_ = "STUB: not implemented"
	return *new(oauth2.AuthCodeOption)
}

func (params Params) TokenRequestOption() oauth2.AuthCodeOption {
	_ = "STUB: not implemented"
	return *new(oauth2.AuthCodeOption)
}

// New returns a parameters supported by the provider.
// You need to pass the code challenge methods defined in RFC7636.
// It returns a zero value if no method is available.
func New(method Method) (Params, error) { _ = "STUB: not implemented"; return *new(Params), nil }

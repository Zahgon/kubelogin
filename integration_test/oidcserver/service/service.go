package service

import (
	"testing"

	"github.com/int128/kubelogin/integration_test/oidcserver/testconfig"
)

func New(t *testing.T, issuerURL string, config testconfig.Config) Service {
	_ = "STUB: not implemented"
	return *new(Service)
}

type service struct {
	config                    testconfig.Config
	t                         *testing.T
	issuerURL                 string
	lastAuthenticationRequest *AuthenticationRequest
	lastTokenResponse         *TokenResponse
}

func (svc *service) IssuerURL() string { _ = "STUB: not implemented"; return "" }

func (svc *service) SetConfig(cfg testconfig.Config) { _ = "STUB: not implemented"; return }

func (svc *service) LastTokenResponse() *TokenResponse { _ = "STUB: not implemented"; return nil }

func (svc *service) Discovery() *DiscoveryResponse {
	_ = "STUB: not implemented"
	// based on https://accounts.google.com/.well-known/openid-configuration
	return nil
}

func (svc *service) GetCertificates() *CertificatesResponse { _ = "STUB: not implemented"; return nil }

func (svc *service) AuthenticateCode(req AuthenticationRequest) (code string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (svc *service) Exchange(req TokenRequest) (*TokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// https://tools.ietf.org/html/rfc7636#section-4.6

func computeS256Challenge(verifier string) string { _ = "STUB: not implemented"; return "" }

func (svc *service) AuthenticatePassword(username, password, scope string) (*TokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (svc *service) Refresh(refreshToken string) (*TokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package kubeconfig

import (
	"testing"
)

// Values represents values in .kubeconfig template.
type Values struct {
	Issuer                      string
	ExtraScopes                 string
	IDPCertificateAuthority     string
	IDPCertificateAuthorityData string
	IDToken                     string
	RefreshToken                string
}

// Create creates a kubeconfig file and returns path to it.
func Create(t *testing.T, v *Values) string { _ = "STUB: not implemented"; return "" }

type AuthProviderConfig struct {
	IDToken      string `yaml:"id-token"`
	RefreshToken string `yaml:"refresh-token"`
}

// Verify returns true if the kubeconfig has valid values.
func Verify(t *testing.T, kubeconfig string, want AuthProviderConfig) {
	_ = "STUB: not implemented"
	return
}

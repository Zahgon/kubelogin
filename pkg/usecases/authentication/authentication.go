package authentication

import (
	"context"

	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/oidc/client"
	"github.com/int128/kubelogin/pkg/tlsclientconfig"
	"github.com/int128/kubelogin/pkg/usecases/authentication/authcode"
	"github.com/int128/kubelogin/pkg/usecases/authentication/clientcredentials"
	"github.com/int128/kubelogin/pkg/usecases/authentication/devicecode"
	"github.com/int128/kubelogin/pkg/usecases/authentication/ropc"
)

// Set provides the use-case of Authentication.
var Set = wire.NewSet(
	wire.Struct(new(Authentication), "*"),
	wire.Bind(new(Interface), new(*Authentication)),
	wire.Struct(new(authcode.Browser), "*"),
	wire.Struct(new(authcode.Keyboard), "*"),
	wire.Struct(new(ropc.ROPC), "*"),
	wire.Struct(new(devicecode.DeviceCode), "*"),
	wire.Struct(new(clientcredentials.ClientCredentials), "*"),
)

type Interface interface {
	Do(ctx context.Context, in Input) (*Output, error)
}

// Input represents an input DTO of the Authentication use-case.
type Input struct {
	Provider        oidc.Provider
	GrantOptionSet  GrantOptionSet
	CachedTokenSet  *oidc.TokenSet // optional
	TLSClientConfig tlsclientconfig.Config
}

type GrantOptionSet struct {
	AuthCodeBrowserOption   *authcode.BrowserOption
	AuthCodeKeyboardOption  *authcode.KeyboardOption
	ROPCOption              *ropc.Option
	DeviceCodeOption        *devicecode.Option
	ClientCredentialsOption *client.GetTokenByClientCredentialsInput
}

// AuthRequestExtraParams returns the extra parameters for the auth request
// from whichever grant option is set.
func (g GrantOptionSet) AuthRequestExtraParams() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Convert map[string][]string back to map[string]string

// Output represents an output DTO of the Authentication use-case.
type Output struct {
	TokenSet oidc.TokenSet
}

// Authentication provides the internal use-case of authentication.
//
// If the IDToken is not set, it performs the authentication flow.
// If the IDToken is valid, it does nothing.
// If the IDToken has expired and the RefreshToken is set, it refreshes the token.
// If the RefreshToken has expired, it performs the authentication flow.
//
// The authentication flow is determined as:
//
// If the Username is not set, it performs the authorization code flow.
// Otherwise, it performs the resource owner password credentials flow.
// If the Password is not set, it asks a password by the prompt.
type Authentication struct {
	ClientFactory     client.FactoryInterface
	Logger            logger.Interface
	AuthCodeBrowser   *authcode.Browser
	AuthCodeKeyboard  *authcode.Keyboard
	ROPC              *ropc.ROPC
	DeviceCode        *devicecode.DeviceCode
	ClientCredentials *clientcredentials.ClientCredentials
}

func (u *Authentication) Do(ctx context.Context, in Input) (*Output, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

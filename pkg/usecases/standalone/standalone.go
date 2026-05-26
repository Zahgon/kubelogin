package standalone

import (
	"context"

	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/infrastructure/clock"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/kubeconfig"
	"github.com/int128/kubelogin/pkg/kubeconfig/loader"
	"github.com/int128/kubelogin/pkg/kubeconfig/writer"
	"github.com/int128/kubelogin/pkg/tlsclientconfig"
	"github.com/int128/kubelogin/pkg/usecases/authentication"
)

// Set provides the use-case.
var Set = wire.NewSet(
	wire.Struct(new(Standalone), "*"),
	wire.Bind(new(Interface), new(*Standalone)),
)

type Interface interface {
	Do(ctx context.Context, in Input) error
}

// Input represents an input DTO of the use-case.
type Input struct {
	KubeconfigFilename string                 // Default to the environment variable or global config as kubectl
	KubeconfigContext  kubeconfig.ContextName // Default to the current context but ignored if KubeconfigUser is set
	KubeconfigUser     kubeconfig.UserName    // Default to the user of the context
	GrantOptionSet     authentication.GrantOptionSet
	TLSClientConfig    tlsclientconfig.Config
}

const oidcConfigErrorMessage = `No configuration found.
You need to set up the OIDC provider, role binding, Kubernetes API server and kubeconfig.
To show the setup instruction:

	kubectl oidc-login setup

See https://github.com/int128/kubelogin for more.
`

// Standalone provides the use case of explicit login.
//
// If the current auth provider is not oidc, show the error.
// If the kubeconfig has a valid token, do nothing.
// Otherwise, update the kubeconfig.
type Standalone struct {
	Authentication   authentication.Interface
	KubeconfigLoader loader.Interface
	KubeconfigWriter writer.Interface
	Logger           logger.Interface
	Clock            clock.Interface
}

func (u *Standalone) Do(ctx context.Context, in Input) error { _ = "STUB: not implemented"; return nil }

// Skip verification of the token to reduce time of a discovery request.
// Here it trusts the signature and claims and checks only expiration,
// because the token has been verified before caching.

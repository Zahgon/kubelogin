// Package credentialplugin provides the use-cases for running as a client-go credentials plugin.
//
// See https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins
package credentialplugin

import (
	"context"

	"github.com/google/wire"
	credentialpluginreader "github.com/int128/kubelogin/pkg/credentialplugin/reader"
	credentialpluginwriter "github.com/int128/kubelogin/pkg/credentialplugin/writer"
	"github.com/int128/kubelogin/pkg/infrastructure/clock"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/tlsclientconfig"
	"github.com/int128/kubelogin/pkg/tokencache"
	"github.com/int128/kubelogin/pkg/tokencache/repository"
	"github.com/int128/kubelogin/pkg/usecases/authentication"
)

var Set = wire.NewSet(
	wire.Struct(new(GetToken), "*"),
	wire.Bind(new(Interface), new(*GetToken)),
)

type Interface interface {
	Do(ctx context.Context, in Input) error
}

// Input represents an input DTO of the GetToken use-case.
type Input struct {
	Provider         oidc.Provider
	ForceRefresh     bool
	TokenCacheConfig tokencache.Config
	GrantOptionSet   authentication.GrantOptionSet
	TLSClientConfig  tlsclientconfig.Config
}

type GetToken struct {
	Authentication         authentication.Interface
	TokenCacheRepository   repository.Interface
	CredentialPluginReader credentialpluginreader.Interface
	CredentialPluginWriter credentialpluginwriter.Interface
	Logger                 logger.Interface
	Clock                  clock.Interface
}

func (u *GetToken) Do(ctx context.Context, in Input) error { _ = "STUB: not implemented"; return nil }

// Skip verification of the token to reduce time of a discovery request.
// Here it trusts the signature and claims and checks only expiration,
// because the token has been verified before caching.

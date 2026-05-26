// Package setup provides the use case of setting up environment.
package setup

import (
	"context"
	"strconv"
	"text/template"

	_ "embed"

	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/tlsclientconfig"
	"github.com/int128/kubelogin/pkg/usecases/authentication"
)

var Set = wire.NewSet(
	wire.Struct(new(Setup), "*"),
	wire.Bind(new(Interface), new(*Setup)),
)

type Interface interface {
	Do(ctx context.Context, in Input) error
}

type Setup struct {
	Authentication authentication.Interface
	Logger         logger.Interface
}

//go:embed setup.md
var setupMarkdown string

var setupTemplate = template.Must(template.New("setup.md").Funcs(template.FuncMap{
	"quote": strconv.Quote,
}).Parse(setupMarkdown))

type Input struct {
	IssuerURL       string
	ClientID        string
	ClientSecret    string
	RedirectURL     string
	ExtraScopes     []string
	UseAccessToken  bool
	RequestHeaders  map[string]string
	PKCEMethod      oidc.PKCEMethod
	GrantOptionSet  authentication.GrantOptionSet
	TLSClientConfig tlsclientconfig.Config
	ChangedFlags    []string
}

func (u Setup) Do(ctx context.Context, in Input) error { _ = "STUB: not implemented"; return nil }

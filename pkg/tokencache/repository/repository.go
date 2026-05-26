package repository

import (
	"io"

	"github.com/google/wire"
	"github.com/int128/kubelogin/pkg/oidc"
	"github.com/int128/kubelogin/pkg/tokencache"
)

// Set provides an implementation and interface for Kubeconfig.
var Set = wire.NewSet(
	wire.Struct(new(Repository), "*"),
	wire.Bind(new(Interface), new(*Repository)),
)

type Interface interface {
	FindByKey(config tokencache.Config, key tokencache.Key) (*oidc.TokenSet, error)
	Save(config tokencache.Config, key tokencache.Key, tokenSet oidc.TokenSet) error
	Lock(config tokencache.Config, key tokencache.Key) (io.Closer, error)
	DeleteAll(config tokencache.Config) error
}

type entity struct {
	IDToken      string `json:"id_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Repository provides access to the token cache on the local filesystem.
// Filename of a token cache is sha256 digest of the issuer, zero-character and client ID.
type Repository struct{}

// keyringService is used to namespace the keyring access.
// Some implementations may also display this string when prompting the user
// for allowing access.
const keyringService = "kubelogin"

// keyringItemPrefix is used as the prefix in the keyring items.
const keyringItemPrefix = "kubelogin/tokencache/"

func (r *Repository) FindByKey(config tokencache.Config, key tokencache.Key) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readFromFile(config tokencache.Config, checksum string) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readFromKeyring(checksum string) (*oidc.TokenSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeKey(b []byte) (*oidc.TokenSet, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Repository) Save(config tokencache.Config, key tokencache.Key, tokenSet oidc.TokenSet) error {
	_ = "STUB: not implemented"
	return nil
}

func writeToFile(config tokencache.Config, checksum string, tokenSet oidc.TokenSet) error {
	_ = "STUB: not implemented"
	return nil
}

func writeToKeyring(checksum string, tokenSet oidc.TokenSet) error {
	_ = "STUB: not implemented"
	return nil
}

// Implement io.Closer for noneStorage type
type noneStorageCloser struct{}

func (c noneStorageCloser) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Repository) Lock(config tokencache.Config, key tokencache.Key) (io.Closer, error) {
	_ = "STUB: not implemented"
	return *new(io.Closer), nil
}

// NOTE: Both keyring and disk storage types use files for locking
// No sensitive data is stored in the lock file

// Do not lock the token cache file.
// https://github.com/int128/kubelogin/issues/1144

func (r *Repository) DeleteAll(config tokencache.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeKey(tokenSet oidc.TokenSet) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func computeChecksum(key tokencache.Key) (string, error) { _ = "STUB: not implemented"; return "", nil }

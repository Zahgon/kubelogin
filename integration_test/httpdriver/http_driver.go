// Package httpdriver provides a test double of the browser.
package httpdriver

import (
	"context"
	"crypto/tls"
	"testing"
)

type Config struct {
	TLSConfig    *tls.Config
	BodyContains string
}

// New returns a client to simulate browser access.
func New(ctx context.Context, t *testing.T, config Config) *client {
	_ = "STUB: not implemented"
	return nil
}

// Zero returns a client which call is not expected.
func Zero(t *testing.T) *zeroClient { _ = "STUB: not implemented"; return nil }

type client struct {
	ctx    context.Context
	t      *testing.T
	config Config
}

func (c *client) Open(url string) error { _ = "STUB: not implemented"; return nil }

func (c *client) OpenCommand(_ context.Context, url, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

type zeroClient struct {
	t *testing.T
}

func (c *zeroClient) Open(url string) error { _ = "STUB: not implemented"; return nil }

func (c *zeroClient) OpenCommand(_ context.Context, url, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

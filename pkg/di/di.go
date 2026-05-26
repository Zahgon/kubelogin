//go:build wireinject
// +build wireinject

// Package di provides dependency injection.
package di

import (
	"github.com/int128/kubelogin/pkg/cmd"
	"github.com/int128/kubelogin/pkg/infrastructure/browser"
	"github.com/int128/kubelogin/pkg/infrastructure/clock"
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/int128/kubelogin/pkg/infrastructure/stdio"
)

// NewCmd returns an instance of infrastructure.Cmd.
func NewCmd() cmd.Interface { _ = "STUB: not implemented"; return *new(cmd.Interface) }

// dependencies for production

// NewCmdForHeadless returns an instance of infrastructure.Cmd for headless testing.
func NewCmdForHeadless(clock.Interface, stdio.Stdin, stdio.Stdout, logger.Interface, browser.Interface) cmd.Interface {
	_ = "STUB: not implemented"

	// use-cases
	return *new(cmd.Interface)
}

// infrastructure

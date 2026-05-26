package logger

import (
	"github.com/int128/kubelogin/pkg/infrastructure/logger"
	"github.com/spf13/pflag"
)

func New(t testingLogger) *Logger { _ = "STUB: not implemented"; return nil }

type testingLogger interface {
	Logf(format string, v ...interface{})
}

// Logger provides logging facility using testing.T.
type Logger struct {
	t        testingLogger
	maxLevel int
}

func (l *Logger) AddFlags(f *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func (l *Logger) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) V(level int) logger.Verbose {
	_ = "STUB: not implemented"
	return *new(logger.Verbose)
}

func (l *Logger) IsEnabled(level int) bool { _ = "STUB: not implemented"; return false }

type verbose struct {
	t     testingLogger
	level int
}

func (v *verbose) Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

type noopVerbose struct{}

func (*noopVerbose) Infof(string, ...interface{}) { _ = "STUB: not implemented"; return }

package transport

import (
	"net/http"

	"github.com/int128/kubelogin/pkg/infrastructure/logger"
)

const (
	levelDumpHeaders = 2
	levelDumpBody    = 3
)

type WithLogging struct {
	Base   http.RoundTripper
	Logger logger.Interface
}

func (t *WithLogging) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

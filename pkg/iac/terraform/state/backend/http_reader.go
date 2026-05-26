package backend

import (
	pkghttp "github.com/snyk/driftctl/pkg/http"

	"io"
	"net/http"
)

const BackendKeyHTTP = "http"
const BackendKeyHTTPS = "https"

type HTTPBackend struct {
	request *http.Request
	client  pkghttp.HTTPClient
	reader  io.ReadCloser
}

func NewHTTPReader(client pkghttp.HTTPClient, rawURL string, opts *Options) (*HTTPBackend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HTTPBackend) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (h *HTTPBackend) Close() error { _ = "STUB: not implemented"; return nil }

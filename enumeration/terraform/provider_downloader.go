package terraform

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-getter"
)

type ProviderDownloaderInterface interface {
	Download(url, path string) error
}

type ProviderDownloader struct {
	httpclient *http.Client
	unzip      getter.ZipDecompressor
	context    context.Context
}

func NewProviderDownloader() *ProviderDownloader { _ = "STUB: not implemented"; return nil }

func (p *ProviderDownloader) Download(url, path string) error {
	_ = "STUB: not implemented"
	return nil
}

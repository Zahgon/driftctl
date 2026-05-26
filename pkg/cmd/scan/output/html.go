package output

import (
	"embed"
	"html/template"

	"github.com/snyk/driftctl/enumeration/alerter"

	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/snyk/driftctl/pkg/analyser"
)

const HTMLOutputType = "html"
const HTMLOutputExample = "html://PATH/TO/FILE.html"

// assets holds our static web content.
//
//go:embed assets/*
var assets embed.FS

type HTML struct {
	path string
}

type HTMLTemplateParams struct {
	IsSync          bool
	ScanDate        string
	Coverage        int
	Summary         analyser.Summary
	Unmanaged       []*resource.Resource
	Deleted         []*resource.Resource
	Alerts          alerter.Alerts
	Stylesheet      template.CSS
	ScanDuration    string
	ProviderName    string
	ProviderVersion string
	LogoSvg         template.HTML
	FaviconBase64   string
}

func NewHTML(path string) *HTML { _ = "STUB: not implemented"; return nil }

func (c *HTML) Write(analysis *analyser.Analysis) error { _ = "STUB: not implemented"; return nil }

func distinctResourceTypes(resources []*resource.Resource) []string {
	_ = "STUB: not implemented"
	return nil
}

func distinctIaCSources(resources []*resource.Resource) []string {
	_ = "STUB: not implemented"
	return nil
}

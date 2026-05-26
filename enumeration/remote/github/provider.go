package github

import (
	"github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/remote/terraform"
)

type GithubTerraformProvider struct {
	*terraform.TerraformProvider
	name    string
	version string
}

type githubConfig struct {
	Token        string
	Owner        string `cty:"owner"`
	Organization string
}

func NewGithubTerraformProvider(version string, progress enumeration.ProgressCounter, configDir string) (*GithubTerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c githubConfig) getDefaultOwner() string { _ = "STUB: not implemented"; return "" }

func (p GithubTerraformProvider) GetConfig() githubConfig {
	_ = "STUB: not implemented"
	return *new(githubConfig)
}

func (p *GithubTerraformProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *GithubTerraformProvider) Version() string { _ = "STUB: not implemented"; return "" }

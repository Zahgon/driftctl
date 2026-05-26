package github

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type GithubTeamEnumerator struct {
	repository GithubRepository
	factory    resource.ResourceFactory
}

func NewGithubTeamEnumerator(repo GithubRepository, factory resource.ResourceFactory) *GithubTeamEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (g *GithubTeamEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (g *GithubTeamEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

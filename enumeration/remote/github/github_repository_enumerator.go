package github

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type GithubRepositoryEnumerator struct {
	repository GithubRepository
	factory    resource.ResourceFactory
}

func NewGithubRepositoryEnumerator(repo GithubRepository, factory resource.ResourceFactory) *GithubRepositoryEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (g *GithubRepositoryEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (g *GithubRepositoryEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

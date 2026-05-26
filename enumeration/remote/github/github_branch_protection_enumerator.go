package github

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type GithubBranchProtectionEnumerator struct {
	repository GithubRepository
	factory    resource.ResourceFactory
}

func NewGithubBranchProtectionEnumerator(repo GithubRepository, factory resource.ResourceFactory) *GithubBranchProtectionEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (g *GithubBranchProtectionEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (g *GithubBranchProtectionEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

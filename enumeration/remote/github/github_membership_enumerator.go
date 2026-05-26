package github

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type GithubMembershipEnumerator struct {
	Membership GithubRepository
	factory    resource.ResourceFactory
}

func NewGithubMembershipEnumerator(repo GithubRepository, factory resource.ResourceFactory) *GithubMembershipEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (g *GithubMembershipEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (g *GithubMembershipEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package github

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type GithubTeamMembershipEnumerator struct {
	repository GithubRepository
	factory    resource.ResourceFactory
}

func NewGithubTeamMembershipEnumerator(repo GithubRepository, factory resource.ResourceFactory) *GithubTeamMembershipEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (g *GithubTeamMembershipEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (g *GithubTeamMembershipEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

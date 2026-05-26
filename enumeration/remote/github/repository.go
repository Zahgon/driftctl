package github

import (
	"context"

	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/shurcooL/githubv4"
)

type GithubRepository interface {
	ListRepositories() ([]string, error)
	ListTeams() ([]Team, error)
	ListMembership() ([]string, error)
	ListTeamMemberships() ([]string, error)
	ListBranchProtection() ([]string, error)
}

type GithubGraphQLClient interface {
	Query(ctx context.Context, q interface{}, variables map[string]interface{}) error
}

type githubRepository struct {
	client GithubGraphQLClient
	ctx    context.Context
	config githubConfig
	cache  cache.Cache
}

func NewGithubRepository(config githubConfig, c cache.Cache) *githubRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *githubRepository) ListRepositories() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type pageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

type listRepoForOrgQuery struct {
	Organization struct {
		Repositories struct {
			Nodes []struct {
				Name string
			}
			PageInfo pageInfo
		} `graphql:"repositories(first: 100, after: $cursor)"`
	} `graphql:"organization(login: $org)"`
}

func (r *githubRepository) listRepoForOrg() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type listRepoForOwnerQuery struct {
	Viewer struct {
		Repositories struct {
			Nodes []struct {
				Name string
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"repositories(first: 100, after: $cursor, ownerAffiliations: OWNER)"`
	}
}

func (r githubRepository) listRepoForOwner() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type listTeamsQuery struct {
	Organization struct {
		Teams struct {
			Nodes []struct {
				DatabaseId int
				Slug       string
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"teams(first: 100, after: $cursor)"`
	} `graphql:"organization(login: $login)"`
}

type Team struct {
	DatabaseId int
	Slug       string
}

func (r githubRepository) ListTeams() ([]Team, error) { _ = "STUB: not implemented"; return nil, nil }

type listMembership struct {
	Organization struct {
		MembersWithRole struct {
			Nodes []struct {
				Login string
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"membersWithRole(first: 100, after: $cursor)"`
	} `graphql:"organization(login: $login)"`
}

func (r *githubRepository) ListMembership() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type listTeamMembershipsQuery struct {
	Organization struct {
		Team struct {
			Members struct {
				Nodes []struct {
					Login string
				}
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"members(first: 100, after: $cursor)"`
		} `graphql:"team(slug: $slug)"`
	} `graphql:"organization(login: $login)"`
}

func (r githubRepository) ListTeamMemberships() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type listBranchProtectionQuery struct {
	Repository struct {
		BranchProtectionRules struct {
			Nodes []struct {
				Id string
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"branchProtectionRules(first: 1, after: $cursor)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

func (r *githubRepository) ListBranchProtection() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

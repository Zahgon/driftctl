package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
)

type ECRRepository interface {
	ListAllRepositories() ([]*ecr.Repository, error)
	GetRepositoryPolicy(*ecr.Repository) (*ecr.GetRepositoryPolicyOutput, error)
}

type ecrRepository struct {
	client ecriface.ECRAPI
	cache  cache.Cache
}

func NewECRRepository(session *session.Session, c cache.Cache) *ecrRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *ecrRepository) ListAllRepositories() ([]*ecr.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ecrRepository) GetRepositoryPolicy(repo *ecr.Repository) (*ecr.GetRepositoryPolicyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

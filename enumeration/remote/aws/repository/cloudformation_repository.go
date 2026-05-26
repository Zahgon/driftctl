package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type CloudformationRepository interface {
	ListAllStacks() ([]*cloudformation.Stack, error)
}

type cloudformationRepository struct {
	client cloudformationiface.CloudFormationAPI
	cache  cache.Cache
}

func NewCloudformationRepository(session *session.Session, c cache.Cache) *cloudformationRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *cloudformationRepository) ListAllStacks() ([]*cloudformation.Stack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

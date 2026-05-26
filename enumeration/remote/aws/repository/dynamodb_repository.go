package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type DynamoDBRepository interface {
	ListAllTables() ([]*string, error)
}

type dynamoDBRepository struct {
	client dynamodbiface.DynamoDBAPI
	cache  cache.Cache
}

func NewDynamoDBRepository(session *session.Session, c cache.Cache) *dynamoDBRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *dynamoDBRepository) ListAllTables() ([]*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

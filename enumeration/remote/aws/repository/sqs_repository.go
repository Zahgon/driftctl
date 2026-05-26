package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type SQSRepository interface {
	ListAllQueues() ([]*string, error)
	GetQueueAttributes(url string) (*sqs.GetQueueAttributesOutput, error)
}

type sqsRepository struct {
	client sqsiface.SQSAPI
	cache  cache.Cache
}

func NewSQSRepository(session *session.Session, c cache.Cache) *sqsRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *sqsRepository) GetQueueAttributes(url string) (*sqs.GetQueueAttributesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *sqsRepository) ListAllQueues() ([]*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

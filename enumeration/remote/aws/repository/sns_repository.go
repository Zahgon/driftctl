package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type SNSRepository interface {
	ListAllTopics() ([]*sns.Topic, error)
	ListAllSubscriptions() ([]*sns.Subscription, error)
}

type snsRepository struct {
	client snsiface.SNSAPI
	cache  cache.Cache
}

func NewSNSRepository(session *session.Session, c cache.Cache) *snsRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *snsRepository) ListAllTopics() ([]*sns.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *snsRepository) ListAllSubscriptions() ([]*sns.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

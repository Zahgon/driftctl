package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type ELBRepository interface {
	ListAllLoadBalancers() ([]*elb.LoadBalancerDescription, error)
}

type elbRepository struct {
	client elbiface.ELBAPI
	cache  cache.Cache
}

func NewELBRepository(session *session.Session, c cache.Cache) *elbRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *elbRepository) ListAllLoadBalancers() ([]*elb.LoadBalancerDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

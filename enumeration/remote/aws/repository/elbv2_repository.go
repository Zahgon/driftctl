package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type ELBV2Repository interface {
	ListAllLoadBalancers() ([]*elbv2.LoadBalancer, error)
	ListAllLoadBalancerListeners(string) ([]*elbv2.Listener, error)
}

type elbv2Repository struct {
	client elbv2iface.ELBV2API
	cache  cache.Cache
}

func NewELBV2Repository(session *session.Session, c cache.Cache) *elbv2Repository {
	_ = "STUB: not implemented"
	return nil
}

func (r *elbv2Repository) ListAllLoadBalancers() ([]*elbv2.LoadBalancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *elbv2Repository) ListAllLoadBalancerListeners(loadBalancerArn string) ([]*elbv2.Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

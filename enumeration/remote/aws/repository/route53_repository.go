package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
)

type Route53Repository interface {
	ListAllHealthChecks() ([]*route53.HealthCheck, error)
	ListAllZones() ([]*route53.HostedZone, error)
	ListRecordsForZone(zoneId string) ([]*route53.ResourceRecordSet, error)
}

type route53Repository struct {
	client route53iface.Route53API
	cache  cache.Cache
}

func NewRoute53Repository(session *session.Session, c cache.Cache) *route53Repository {
	_ = "STUB: not implemented"
	return nil
}

func (r *route53Repository) ListAllHealthChecks() ([]*route53.HealthCheck, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *route53Repository) ListAllZones() ([]*route53.HostedZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *route53Repository) ListRecordsForZone(zoneId string) ([]*route53.ResourceRecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

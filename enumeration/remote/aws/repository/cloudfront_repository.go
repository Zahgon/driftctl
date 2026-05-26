package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type CloudfrontRepository interface {
	ListAllDistributions() ([]*cloudfront.DistributionSummary, error)
}

type cloudfrontRepository struct {
	client cloudfrontiface.CloudFrontAPI
	cache  cache.Cache
}

func NewCloudfrontRepository(session *session.Session, c cache.Cache) *cloudfrontRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *cloudfrontRepository) ListAllDistributions() ([]*cloudfront.DistributionSummary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

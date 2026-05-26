package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type CloudtrailRepository interface {
	ListAllTrails() ([]*cloudtrail.TrailInfo, error)
}

type cloudtrailRepository struct {
	client cloudtrailiface.CloudTrailAPI
	cache  cache.Cache
}

func NewCloudtrailRepository(session *session.Session, c cache.Cache) *cloudtrailRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *cloudtrailRepository) ListAllTrails() ([]*cloudtrail.TrailInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

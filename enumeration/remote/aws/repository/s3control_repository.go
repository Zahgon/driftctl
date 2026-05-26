package repository

import (
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/snyk/driftctl/enumeration/remote/aws/client"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type S3ControlRepository interface {
	DescribeAccountPublicAccessBlock(accountID string) (*s3control.PublicAccessBlockConfiguration, error)
}

type s3ControlRepository struct {
	clientFactory client.AwsClientFactoryInterface
	cache         cache.Cache
}

func NewS3ControlRepository(factory client.AwsClientFactoryInterface, c cache.Cache) *s3ControlRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *s3ControlRepository) DescribeAccountPublicAccessBlock(accountID string) (*s3control.PublicAccessBlockConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3ControlRepository) shouldSuppressError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

// do not throw the error up if there is no access block config

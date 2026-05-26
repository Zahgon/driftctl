package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/client"
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Repository interface {
	ListAllBuckets() ([]*s3.Bucket, error)
	GetBucketNotification(bucketName, region string) (*s3.NotificationConfiguration, error)
	GetBucketPolicy(bucketName, region string) (*string, error)
	GetBucketPublicAccessBlock(bucketName, region string) (*s3.PublicAccessBlockConfiguration, error)
	ListBucketInventoryConfigurations(bucket *s3.Bucket, region string) ([]*s3.InventoryConfiguration, error)
	ListBucketMetricsConfigurations(bucket *s3.Bucket, region string) ([]*s3.MetricsConfiguration, error)
	ListBucketAnalyticsConfigurations(bucket *s3.Bucket, region string) ([]*s3.AnalyticsConfiguration, error)
	GetBucketLocation(bucketName string) (string, error)
}

type s3Repository struct {
	clientFactory client.AwsClientFactoryInterface
	cache         cache.Cache
}

func NewS3Repository(factory client.AwsClientFactoryInterface, c cache.Cache) *s3Repository {
	_ = "STUB: not implemented"
	return nil
}

func (s *s3Repository) ListAllBuckets() ([]*s3.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3Repository) GetBucketPolicy(bucketName, region string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3Repository) GetBucketPublicAccessBlock(bucketName, region string) (*s3.PublicAccessBlockConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3Repository) GetBucketNotification(bucketName, region string) (*s3.NotificationConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3Repository) notificationIsEmpty(notification *s3.NotificationConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *s3Repository) ListBucketInventoryConfigurations(bucket *s3.Bucket, region string) ([]*s3.InventoryConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3Repository) ListBucketMetricsConfigurations(bucket *s3.Bucket, region string) ([]*s3.MetricsConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3Repository) ListBucketAnalyticsConfigurations(bucket *s3.Bucket, region string) ([]*s3.AnalyticsConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *s3Repository) GetBucketLocation(bucketName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Buckets in Region us-east-1 have a LocationConstraint of null.
// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLocation.html#API_GetBucketLocation_ResponseSyntax

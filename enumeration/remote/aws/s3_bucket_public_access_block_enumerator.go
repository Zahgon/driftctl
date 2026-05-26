package aws

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	tf "github.com/snyk/driftctl/enumeration/remote/terraform"
	"github.com/snyk/driftctl/enumeration/resource"
)

type S3BucketPublicAccessBlockEnumerator struct {
	repository     repository.S3Repository
	factory        resource.ResourceFactory
	providerConfig tf.TerraformProviderConfig
	alerter        alerter.AlerterInterface
}

func NewS3BucketPublicAccessBlockEnumerator(repo repository.S3Repository, factory resource.ResourceFactory, providerConfig tf.TerraformProviderConfig, alerter alerter.AlerterInterface) *S3BucketPublicAccessBlockEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *S3BucketPublicAccessBlockEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *S3BucketPublicAccessBlockEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type CloudfrontDistributionEnumerator struct {
	repository repository.CloudfrontRepository
	factory    resource.ResourceFactory
}

func NewCloudfrontDistributionEnumerator(repo repository.CloudfrontRepository, factory resource.ResourceFactory) *CloudfrontDistributionEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *CloudfrontDistributionEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *CloudfrontDistributionEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

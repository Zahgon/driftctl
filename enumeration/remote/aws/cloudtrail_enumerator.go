package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type CloudtrailEnumerator struct {
	repository repository.CloudtrailRepository
	factory    resource.ResourceFactory
}

func NewCloudtrailEnumerator(repo repository.CloudtrailRepository, factory resource.ResourceFactory) *CloudtrailEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *CloudtrailEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *CloudtrailEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

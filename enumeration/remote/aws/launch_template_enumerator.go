package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type LaunchTemplateEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewLaunchTemplateEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *LaunchTemplateEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *LaunchTemplateEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *LaunchTemplateEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

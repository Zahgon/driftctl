package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type LaunchConfigurationEnumerator struct {
	repository repository.AutoScalingRepository
	factory    resource.ResourceFactory
}

func NewLaunchConfigurationEnumerator(repo repository.AutoScalingRepository, factory resource.ResourceFactory) *LaunchConfigurationEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *LaunchConfigurationEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *LaunchConfigurationEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

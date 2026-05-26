package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AppAutoscalingTargetEnumerator struct {
	repository repository.AppAutoScalingRepository
	factory    resource.ResourceFactory
}

func NewAppAutoscalingTargetEnumerator(repository repository.AppAutoScalingRepository, factory resource.ResourceFactory) *AppAutoscalingTargetEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AppAutoscalingTargetEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AppAutoscalingTargetEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

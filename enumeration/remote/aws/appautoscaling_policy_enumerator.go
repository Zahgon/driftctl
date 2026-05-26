package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AppAutoscalingPolicyEnumerator struct {
	repository repository.AppAutoScalingRepository
	factory    resource.ResourceFactory
}

func NewAppAutoscalingPolicyEnumerator(repository repository.AppAutoScalingRepository, factory resource.ResourceFactory) *AppAutoscalingPolicyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AppAutoscalingPolicyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AppAutoscalingPolicyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

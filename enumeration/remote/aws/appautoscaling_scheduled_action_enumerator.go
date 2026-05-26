package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type AppAutoscalingScheduledActionEnumerator struct {
	repository repository.AppAutoScalingRepository
	factory    resource.ResourceFactory
}

func NewAppAutoscalingScheduledActionEnumerator(repository repository.AppAutoScalingRepository, factory resource.ResourceFactory) *AppAutoscalingScheduledActionEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AppAutoscalingScheduledActionEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AppAutoscalingScheduledActionEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

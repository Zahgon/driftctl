package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type SQSQueuePolicyEnumerator struct {
	repository repository.SQSRepository
	factory    resource.ResourceFactory
}

func NewSQSQueuePolicyEnumerator(repo repository.SQSRepository, factory resource.ResourceFactory) *SQSQueuePolicyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *SQSQueuePolicyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *SQSQueuePolicyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

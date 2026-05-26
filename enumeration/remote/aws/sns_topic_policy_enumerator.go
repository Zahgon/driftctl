package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type SNSTopicPolicyEnumerator struct {
	repository repository.SNSRepository
	factory    resource.ResourceFactory
}

func NewSNSTopicPolicyEnumerator(repo repository.SNSRepository, factory resource.ResourceFactory) *SNSTopicPolicyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *SNSTopicPolicyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *SNSTopicPolicyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

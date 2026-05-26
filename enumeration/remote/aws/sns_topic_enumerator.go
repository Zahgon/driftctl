package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type SNSTopicEnumerator struct {
	repository repository.SNSRepository
	factory    resource.ResourceFactory
}

func NewSNSTopicEnumerator(repo repository.SNSRepository, factory resource.ResourceFactory) *SNSTopicEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *SNSTopicEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *SNSTopicEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ECRRepositoryPolicyEnumerator struct {
	repository repository.ECRRepository
	factory    resource.ResourceFactory
}

func NewECRRepositoryPolicyEnumerator(repo repository.ECRRepository, factory resource.ResourceFactory) *ECRRepositoryPolicyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ECRRepositoryPolicyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ECRRepositoryPolicyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

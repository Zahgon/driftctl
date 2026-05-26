package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ECRRepositoryEnumerator struct {
	repository repository.ECRRepository
	factory    resource.ResourceFactory
}

func NewECRRepositoryEnumerator(repo repository.ECRRepository, factory resource.ResourceFactory) *ECRRepositoryEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ECRRepositoryEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ECRRepositoryEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

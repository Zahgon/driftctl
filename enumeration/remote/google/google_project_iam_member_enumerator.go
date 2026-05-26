package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleProjectIamMemberEnumerator struct {
	repository repository.CloudResourceManagerRepository
	factory    resource.ResourceFactory
}

func NewGoogleProjectIamMemberEnumerator(repo repository.CloudResourceManagerRepository, factory resource.ResourceFactory) *GoogleProjectIamMemberEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleProjectIamMemberEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleProjectIamMemberEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

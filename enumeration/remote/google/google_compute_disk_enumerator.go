package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeDiskEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeDiskEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeDiskEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeDiskEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeDiskEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

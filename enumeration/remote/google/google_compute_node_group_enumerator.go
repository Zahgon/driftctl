package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeNodeGroupEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeNodeGroupEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeNodeGroupEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeNodeGroupEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeNodeGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

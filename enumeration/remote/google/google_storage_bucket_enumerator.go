package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleStorageBucketEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleStorageBucketEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleStorageBucketEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleStorageBucketEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleStorageBucketEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

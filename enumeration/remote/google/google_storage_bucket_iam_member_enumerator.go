package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleStorageBucketIamMemberEnumerator struct {
	repository        repository.AssetRepository
	storageRepository repository.StorageRepository
	factory           resource.ResourceFactory
}

func NewGoogleStorageBucketIamMemberEnumerator(repo repository.AssetRepository, storageRepo repository.StorageRepository, factory resource.ResourceFactory) *GoogleStorageBucketIamMemberEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleStorageBucketIamMemberEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleStorageBucketIamMemberEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

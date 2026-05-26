package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleBigTableInstanceEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleBigTableInstanceEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleBigTableInstanceEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleBigTableInstanceEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleBigTableInstanceEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

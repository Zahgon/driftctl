package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeSslCertificateEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeSslCertificateEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeSslCertificateEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeSslCertificateEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeSslCertificateEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

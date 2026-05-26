package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleDNSManagedZoneEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleDNSManagedZoneEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleDNSManagedZoneEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleDNSManagedZoneEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleDNSManagedZoneEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We should have ID = "projects/cloudskiff-dev-elie/managedZones/example-zone"
// We have projects/cloudskiff-dev-elie/managedZones/2435093289230056557

package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeFirewallEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeFirewallEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeFirewallEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeFirewallEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeFirewallEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

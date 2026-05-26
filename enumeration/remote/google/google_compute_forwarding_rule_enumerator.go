package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeForwardingRuleEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeForwardingRuleEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeForwardingRuleEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeForwardingRuleEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeForwardingRuleEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

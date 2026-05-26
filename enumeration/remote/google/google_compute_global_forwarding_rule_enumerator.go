package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeGlobalForwardingRuleEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeGlobalForwardingRuleEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeGlobalForwardingRuleEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeGlobalForwardingRuleEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeGlobalForwardingRuleEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

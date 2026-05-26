package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2NetworkACLRuleEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2NetworkACLRuleEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2NetworkACLRuleEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2NetworkACLRuleEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2NetworkACLRuleEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Used in default middleware
// Used in default middleware
// Used in default middleware

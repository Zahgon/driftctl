package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

var iamRoleExclusionList = map[string]struct{}{
	// Enabled by default for aws to enable support, not removable
	"AWSServiceRoleForSupport": {},
	// Enabled and not removable for every org account
	"AWSServiceRoleForOrganizations": {},
	// Not manageable by IaC and set by default
	"AWSServiceRoleForTrustedAdvisor": {},
}

type IamRoleEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamRoleEnumerator(repository repository.IAMRepository, factory resource.ResourceFactory) *IamRoleEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamRoleEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func awsIamRoleShouldBeIgnored(roleName string) bool { _ = "STUB: not implemented"; return false }

func (e *IamRoleEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

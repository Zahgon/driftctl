package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Default network acl rules should not be shown as unmanaged as they are present by default
// This middleware ignores default network acl rules from unmanaged resources if they are not managed by IaC
type AwsDefaultNetworkACLRule struct{}

func NewAwsDefaultNetworkACLRule() AwsDefaultNetworkACLRule {
	_ = "STUB: not implemented"
	return *new(AwsDefaultNetworkACLRule)
}

func (m AwsDefaultNetworkACLRule) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than ACL rules

// Ignore non default ACL rules

// Check if resource is managed by IaC

// Include resource if it's managed in IaC

// Else, resource is not added to newRemoteResources slice so it will be ignored

func (m *AwsDefaultNetworkACLRule) isDefaultACLRule(res *resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

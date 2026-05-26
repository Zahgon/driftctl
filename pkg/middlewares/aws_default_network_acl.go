package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Default network ACL should not be shown as unmanaged as they are present by default
// This middleware ignores default network ACL from unmanaged resources if they are not managed by IaC
type AwsDefaultNetworkACL struct{}

func NewAwsDefaultNetworkACL() AwsDefaultNetworkACL {
	_ = "STUB: not implemented"
	return *new(AwsDefaultNetworkACL)
}

func (m AwsDefaultNetworkACL) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than network ACLs

// Check if resource is managed by IaC

// Include resource if it's managed in IaC

// Else, resource is not added to newRemoteResources slice so it will be ignored

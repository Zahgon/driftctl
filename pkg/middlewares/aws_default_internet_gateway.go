package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Each default vpc has an internet gateway attached that should not be seen as unmanaged if not managed by IaC
// This middleware ignores default internet gateway from unmanaged resources if not managed by IaC
type AwsDefaultInternetGateway struct{}

func NewAwsDefaultInternetGateway() AwsDefaultInternetGateway {
	_ = "STUB: not implemented"
	return *new(AwsDefaultInternetGateway)
}

func (m AwsDefaultInternetGateway) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than internet gateways

// Ignore all non-default internet gateways

// Check if internet gateway is managed by IaC

// Include resource if it's managed in IaC

// Else, resource is not added to newRemoteResources slice so it will be ignored

// Return true if the internet gateway is the default one (e.g. attached to the default vpc)
func isDefaultInternetGateway(internetGateway *resource.Resource, remoteResources *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

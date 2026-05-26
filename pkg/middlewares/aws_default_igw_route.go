package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Each region has a default vpc which has an internet gateway attached and thus the route table of this
// same vpc has a default route (0.0.0.0/0) that should not be seen as unmanaged if not managed by IaC
// This middleware ignores the above route from unmanaged resources if not managed by IaC
type AwsDefaultInternetGatewayRoute struct{}

func NewAwsDefaultInternetGatewayRoute() AwsDefaultInternetGatewayRoute {
	_ = "STUB: not implemented"
	return *new(AwsDefaultInternetGatewayRoute)
}

func (m AwsDefaultInternetGatewayRoute) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than routes

// Ignore all routes except the one that came from the default internet gateway

// Check if route is managed by IaC

// Include resource if it's managed in IaC

// Else, resource is not added to newRemoteResources slice so it will be ignored

// Return true if the route's target is the default internet gateway (e.g. attached to the default vpc)
func isDefaultInternetGatewayRoute(route *resource.Resource, remoteResources *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Default routes should not be shown as unmanaged as they are present by default
// This middleware ignores default routes from unmanaged resources if they are not managed by IaC
type AwsDefaultRoute struct{}

func NewAwsDefaultRoute() AwsDefaultRoute { _ = "STUB: not implemented"; return *new(AwsDefaultRoute) }

func (m AwsDefaultRoute) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than routes

// Ignore all non-default routes, check if route is coming from table creation

// Check if route is managed by IaC

// Include resource if it's managed in IaC

// Else, resource is not added to newRemoteResources slice so it will be ignored

// "route": route.String(), TODO

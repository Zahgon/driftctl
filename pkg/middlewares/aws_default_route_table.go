package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Default route table should not be shown as unmanaged as they are present by default
// This middleware ignores default route table from unmanaged resources if they are not managed by IaC
type AwsDefaultRouteTable struct{}

func NewAwsDefaultRouteTable() AwsDefaultRouteTable {
	_ = "STUB: not implemented"
	return *new(AwsDefaultRouteTable)
}

func (m AwsDefaultRouteTable) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than default RouteTable

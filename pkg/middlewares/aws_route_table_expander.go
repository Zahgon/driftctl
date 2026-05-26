package middlewares

import (
	"github.com/snyk/driftctl/enumeration/alerter"

	"github.com/snyk/driftctl/enumeration/resource"
)

// Explodes routes found in aws_default_route_table.route and aws_route_table.route to dedicated resources
type AwsRouteTableExpander struct {
	alerter         alerter.AlerterInterface
	resourceFactory resource.ResourceFactory
}

func NewAwsRouteTableExpander(alerter alerter.AlerterInterface, resourceFactory resource.ResourceFactory) AwsRouteTableExpander {
	_ = "STUB: not implemented"
	return *new(AwsRouteTableExpander)
}

func (m AwsRouteTableExpander) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than (default) routes tables

func (m *AwsRouteTableExpander) handleTable(table *resource.Resource, results *[]*resource.Resource, resourcesFromState []*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Don't expand if the route already exists as a dedicated resource

func (m *AwsRouteTableExpander) handleDefaultTable(table *resource.Resource, results *[]*resource.Resource, resourcesFromState []*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Don't expand if the route already exists as a dedicated resource

func (m *AwsRouteTableExpander) routeExists(routeId string, resourcesFromState []*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

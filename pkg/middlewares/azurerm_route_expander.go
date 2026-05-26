package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Explodes routes found in azurerm_route_table.route from state resources to dedicated resources
type AzurermRouteExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewAzurermRouteExpander(resourceFactory resource.ResourceFactory) AzurermRouteExpander {
	_ = "STUB: not implemented"
	return *new(AzurermRouteExpander)
}

func (m AzurermRouteExpander) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than route tables

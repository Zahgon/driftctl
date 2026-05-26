package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Explodes subnet found in azurerm_virtual_network.subnet from state resources to dedicated resources
type AzurermSubnetExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewAzurermSubnetExpander(resourceFactory resource.ResourceFactory) AzurermSubnetExpander {
	_ = "STUB: not implemented"
	return *new(AzurermSubnetExpander)
}

func (m AzurermSubnetExpander) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than azurerm_virtual_network

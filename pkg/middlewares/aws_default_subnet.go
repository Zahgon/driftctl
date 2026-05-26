package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Default subnet should not be shown as unmanaged as they are present by default
// This middleware ignores default subnet from unmanaged resources if they are not managed by IaC
type AwsDefaultSubnet struct{}

func NewAwsDefaultSubnet() AwsDefaultSubnet {
	_ = "STUB: not implemented"
	return *new(AwsDefaultSubnet)
}

func (m AwsDefaultSubnet) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than default Subnet

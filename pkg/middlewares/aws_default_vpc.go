package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Default VPC should not be shown as unmanaged as they are present by default
// This middleware ignores default VPC from unmanaged resources if they are not managed by IaC
type AwsDefaultVPC struct{}

func NewAwsDefaultVPC() AwsDefaultVPC { _ = "STUB: not implemented"; return *new(AwsDefaultVPC) }

func (m AwsDefaultVPC) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than default VPC

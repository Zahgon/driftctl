package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Remove default security group from remote resources
type VPCDefaultSecurityGroupSanitizer struct{}

func NewVPCDefaultSecurityGroupSanitizer() VPCDefaultSecurityGroupSanitizer {
	_ = "STUB: not implemented"
	return *new(VPCDefaultSecurityGroupSanitizer)
}

func (m VPCDefaultSecurityGroupSanitizer) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than default security group

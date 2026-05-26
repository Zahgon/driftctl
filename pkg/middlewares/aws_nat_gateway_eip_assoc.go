package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type AwsNatGatewayEipAssoc struct{}

func NewAwsNatGatewayEipAssoc() AwsNatGatewayEipAssoc {
	_ = "STUB: not implemented"
	return *new(AwsNatGatewayEipAssoc)
}

// When creating a nat gateway, we associate an EIP to the gateway
// It implies that driftctl read a aws_eip_association resource from remote
// As we cannot use aws_eip_association in terraform to assign an eip to an aws_nat_gateway
// we should remove this association to ensure we do not output noise in unmanaged resources
func (a AwsNatGatewayEipAssoc) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than aws_eip_association

// Ignore all resources other than aws_eip_association

func (a AwsNatGatewayEipAssoc) isAssociatedToNatGateway(cur *resource.Resource, resourceSet *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	// Search for a nat gateway associated with our EIP
	return false
}

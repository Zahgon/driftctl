package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Used to reconcile API Gateway domain names (v1 and v2) from both remote
// and state resources because v1|v2 AWS SDK list endpoints return all domain
// names without distinction
type AwsApiGatewayDomainNamesReconciler struct{}

func NewAwsApiGatewayDomainNamesReconciler() AwsApiGatewayDomainNamesReconciler {
	_ = "STUB: not implemented"
	return *new(AwsApiGatewayDomainNamesReconciler)
}

func (m AwsApiGatewayDomainNamesReconciler) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than aws_api_gateway_domain_name and aws_apigatewayv2_domain_name

// Find a matching state resources

// Keep track of the resource if it's managed in IaC

// If we're here, it means that we are left with unmanaged domain names
// in both v1 and v2 format. Let's group real and duplicate domain names
// in a slice

// We only want to show to our end users unmanaged v1 domain names
// To do that, we're going to loop over unmanaged domain names to delete duplicates
// and leave after that only v1 domain names (e.g. remove v2 ones)

// Remove duplicates (e.g. same id, the opposite type)

// Now keep only v1 domain names

// Finally, add both managed and unmanaged resources to remote resources

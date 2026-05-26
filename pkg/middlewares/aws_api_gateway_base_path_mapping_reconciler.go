package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// AwsApiGatewayBasePathMappingReconciler is used to reconcile API Gateway base path mapping (v1 and v2)
// from both remote and state resources because v1|v2 AWS SDK list endpoints return all mappings
// without distinction.
type AwsApiGatewayBasePathMappingReconciler struct{}

func NewAwsApiGatewayBasePathMappingReconciler() AwsApiGatewayBasePathMappingReconciler {
	_ = "STUB: not implemented"
	return *new(AwsApiGatewayBasePathMappingReconciler)
}

func (m AwsApiGatewayBasePathMappingReconciler) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than aws_api_gateway_base_path_mapping and aws_apigatewayv2_api_mapping

// Find a matching state resources

// Keep track of the resource if it's managed in IaC

// If we're here, it means that we are left with unmanaged path mappings
// in both v1 and v2 format. Let's group real and duplicate path mappings
// in a slice

// We only want to show to our end users unmanaged v1 path mappings
// To do that, we're going to loop over unmanaged path mappings to delete duplicates
// and leave after that only v1 path mappings (e.g. remove v2 ones)

// Remove duplicates (e.g. same id, the opposite type)

// Now keep only v1 path mappings

// Finally, add both managed and unmanaged resources to remote resources

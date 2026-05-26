package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Explodes policy found in aws_api_gateway_rest_api.policy from state resources to dedicated resources
type AwsApiGatewayRestApiPolicyExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsApiGatewayRestApiPolicyExpander(resourceFactory resource.ResourceFactory) AwsApiGatewayRestApiPolicyExpander {
	_ = "STUB: not implemented"
	return *new(AwsApiGatewayRestApiPolicyExpander)
}

func (m AwsApiGatewayRestApiPolicyExpander) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than api_gateway_rest_api

func (m *AwsApiGatewayRestApiPolicyExpander) handlePolicy(api *resource.Resource, results *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Return true if the rest api has a aws_api_gateway_rest_api_policy resource attached to itself.
// It is mandatory since it's possible to have a aws_api_gateway_rest_api with an inline policy
// AND a aws_api_gateway_rest_api_policy resource at the same time. At the end, on the AWS console,
// the aws_api_gateway_rest_api_policy will be used.
func hasRestApiPolicyAttached(api string, resourcesFromState *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

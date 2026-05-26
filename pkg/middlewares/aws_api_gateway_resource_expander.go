package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Explodes api gateway default resource found in aws_api_gateway_rest_api.root_resource_id from state resources to dedicated resources
type AwsApiGatewayResourceExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsApiGatewayResourceExpander(resourceFactory resource.ResourceFactory) AwsApiGatewayResourceExpander {
	_ = "STUB: not implemented"
	return *new(AwsApiGatewayResourceExpander)
}

func (m AwsApiGatewayResourceExpander) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than aws_api_gateway_rest_api

func (m *AwsApiGatewayResourceExpander) handleResource(api *resource.Resource, results *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

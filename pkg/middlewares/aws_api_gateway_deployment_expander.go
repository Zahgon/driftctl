package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Create a aws_api_gateway_stage resource from a aws_api_gateway_deployment resource and ignore the latter resource
// since we don't support it
type AwsApiGatewayDeploymentExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsApiGatewayDeploymentExpander(resourceFactory resource.ResourceFactory) AwsApiGatewayDeploymentExpander {
	_ = "STUB: not implemented"
	return *new(AwsApiGatewayDeploymentExpander)
}

func (m AwsApiGatewayDeploymentExpander) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

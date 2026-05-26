package middlewares

import (
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/snyk/driftctl/enumeration/resource"
)

// Explodes the body attribute of api gateway apis v1|v2 to dedicated resources as per Terraform documentations
// (https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_rest_api)
// (https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/apigatewayv2_api)
type AwsApiGatewayApiExpander struct {
	resourceFactory resource.ResourceFactory
}

type OpenAPIAwsExtensions struct {
	GatewayResponses map[string]interface{} `json:"x-amazon-apigateway-gateway-responses"`
}

type OpenAPIAwsMethodExtensions struct {
	Integration map[string]interface{} `json:"x-amazon-apigateway-integration"`
}

func NewAwsApiGatewayApiExpander(resourceFactory resource.ResourceFactory) AwsApiGatewayApiExpander {
	_ = "STUB: not implemented"
	return *new(AwsApiGatewayApiExpander)
}

func (m AwsApiGatewayApiExpander) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than aws_api_gateway_rest_api && aws_apigatewayv2_api

func (m *AwsApiGatewayApiExpander) handleBody(api *resource.Resource, results, remoteResources *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// It's an OpenAPI v3 document

// It's an OpenAPI v2 document

func (m *AwsApiGatewayApiExpander) handleBodyOpenAPIv3(api *resource.Resource, doc *openapi3.T, results, remoteResources *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *AwsApiGatewayApiExpander) handleBodyOpenAPIv3GatewayV2(api *resource.Resource, doc *openapi3.T, results, remoteResources *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// The types are similar structurally between the openapi2 and openapi3
// libraries, but without generics we can't really de-dup this witout code
// generation, which isn't worth it for this short function.
func (m *AwsApiGatewayApiExpander) handleBodyOpenAPIv2GatewayV2(api *resource.Resource, doc *openapi2.T, results, remoteResources *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func findMatchingOpenAPIDerivedRoute(desiredApiID, desiredPath, desiredMethod string, remoteResources *[]*resource.Resource) *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

func findMatchingOpenAPIDerivedIntegration(desiredApiID string, desiredIntegration *OpenAPIAwsMethodExtensions, remoteResources *[]*resource.Resource) *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

// This is nilable in MOCK type only, and they cannot be embedded

func (m *AwsApiGatewayApiExpander) handleBodyOpenAPIv2(api *resource.Resource, doc *openapi2.T, results, remoteResources *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Create resources based on our OpenAPIAwsExtensions struct
func (m *AwsApiGatewayApiExpander) createExtensionsResources(apiId string, extensions map[string]interface{}, results *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Create resources based on our OpenAPIAwsMethodExtensions struct
func (m *AwsApiGatewayApiExpander) createMethodExtensionsResources(apiId, resourceId, httpMethod string, extensions map[string]interface{}, results *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Create aws_api_gateway_resource resource
func (m *AwsApiGatewayApiExpander) createApiGatewayResource(apiId, path string, results, remoteResources *[]*resource.Resource) *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

// Create aws_api_gateway_gateway_response resource
func (m *AwsApiGatewayApiExpander) createApiGatewayGatewayResponse(apiId, gtwResponse string, results *[]*resource.Resource) {
	_ = "STUB: not implemented"
	return
}

// Returns the aws_api_gateway_resource resource that matches the path attribute
func foundMatchingResource(apiId, path string, remoteResources *[]*resource.Resource) *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

// Create aws_api_gateway_method resource
func (m *AwsApiGatewayApiExpander) createApiGatewayMethod(apiId, resourceId, httpMethod string, results *[]*resource.Resource) {
	_ = "STUB: not implemented"
	return
}

// Create aws_api_gateway_method_response resource
func (m *AwsApiGatewayApiExpander) createApiGatewayMethodResponse(apiId, resourceId, httpMethod, statusCode string, results *[]*resource.Resource) {
	_ = "STUB: not implemented"
	return
}

// Decode openapi.Extensions into our custom OpenAPIAwsExtensions struct that follows AWS
// OpenAPI addons.
func decodeExtensions(extensions map[string]interface{}) (*OpenAPIAwsExtensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create aws_api_gateway_integration resource
func (m *AwsApiGatewayApiExpander) createApiGatewayIntegration(apiId, resourceId, httpMethod string, results *[]*resource.Resource) {
	_ = "STUB: not implemented"
	return
}

// Create aws_api_gateway_integration resource
func (m *AwsApiGatewayApiExpander) createApiGatewayIntegrationResponse(apiId, resourceId, httpMethod, statusCode string, results *[]*resource.Resource) {
	_ = "STUB: not implemented"
	return
}

// Decode openapi.Method.Extensions into our custom OpenAPIAwsMethodExtensions struct that follows AWS
// OpenAPI addons.
func decodeMethodExtensions(extensions map[string]interface{}) (*OpenAPIAwsMethodExtensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

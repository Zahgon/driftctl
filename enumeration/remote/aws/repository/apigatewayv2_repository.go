package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
)

type ApiGatewayV2Repository interface {
	ListAllApis() ([]*apigatewayv2.Api, error)
	ListAllApiRoutes(apiId *string) ([]*apigatewayv2.Route, error)
	ListAllApiDeployments(apiId *string) ([]*apigatewayv2.Deployment, error)
	ListAllVpcLinks() ([]*apigatewayv2.VpcLink, error)
	ListAllApiAuthorizers(string) ([]*apigatewayv2.Authorizer, error)
	ListAllApiIntegrations(string) ([]*apigatewayv2.Integration, error)
	ListAllApiModels(string) ([]*apigatewayv2.Model, error)
	ListAllApiStages(string) ([]*apigatewayv2.Stage, error)
	ListAllApiRouteResponses(string, string) ([]*apigatewayv2.RouteResponse, error)
	ListAllApiMappings(string) ([]*apigatewayv2.ApiMapping, error)
	ListAllApiIntegrationResponses(string, string) ([]*apigatewayv2.IntegrationResponse, error)
}
type apigatewayv2Repository struct {
	client apigatewayv2iface.ApiGatewayV2API
	cache  cache.Cache
}

func NewApiGatewayV2Repository(session *session.Session, c cache.Cache) *apigatewayv2Repository {
	_ = "STUB: not implemented"
	return nil
}

func (r *apigatewayv2Repository) ListAllApis() ([]*apigatewayv2.Api, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiRoutes(apiID *string) ([]*apigatewayv2.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiDeployments(apiID *string) ([]*apigatewayv2.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllVpcLinks() ([]*apigatewayv2.VpcLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiAuthorizers(apiId string) ([]*apigatewayv2.Authorizer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiIntegrations(apiId string) ([]*apigatewayv2.Integration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiModels(apiId string) ([]*apigatewayv2.Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiStages(apiId string) ([]*apigatewayv2.Stage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiIntegrationResponses(apiId, integrationId string) ([]*apigatewayv2.IntegrationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiRouteResponses(apiId, routeId string) ([]*apigatewayv2.RouteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayv2Repository) ListAllApiMappings(domainName string) ([]*apigatewayv2.ApiMapping, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

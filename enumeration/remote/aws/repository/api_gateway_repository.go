package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
)

type ApiGatewayRepository interface {
	ListAllRestApis() ([]*apigateway.RestApi, error)
	GetAccount() (*apigateway.Account, error)
	ListAllApiKeys() ([]*apigateway.ApiKey, error)
	ListAllRestApiAuthorizers(string) ([]*apigateway.Authorizer, error)
	ListAllRestApiStages(string) ([]*apigateway.Stage, error)
	ListAllRestApiResources(string) ([]*apigateway.Resource, error)
	ListAllDomainNames() ([]*apigateway.DomainName, error)
	ListAllVpcLinks() ([]*apigateway.UpdateVpcLinkOutput, error)
	ListAllRestApiRequestValidators(string) ([]*apigateway.UpdateRequestValidatorOutput, error)
	ListAllDomainNameBasePathMappings(string) ([]*apigateway.BasePathMapping, error)
	ListAllRestApiModels(string) ([]*apigateway.Model, error)
	ListAllRestApiGatewayResponses(string) ([]*apigateway.UpdateGatewayResponseOutput, error)
}

type apigatewayRepository struct {
	client apigatewayiface.APIGatewayAPI
	cache  cache.Cache
}

func NewApiGatewayRepository(session *session.Session, c cache.Cache) *apigatewayRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *apigatewayRepository) ListAllRestApis() ([]*apigateway.RestApi, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) GetAccount() (*apigateway.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllApiKeys() ([]*apigateway.ApiKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllRestApiAuthorizers(apiId string) ([]*apigateway.Authorizer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllRestApiStages(apiId string) ([]*apigateway.Stage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllRestApiResources(apiId string) ([]*apigateway.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllDomainNames() ([]*apigateway.DomainName, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllVpcLinks() ([]*apigateway.UpdateVpcLinkOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllRestApiRequestValidators(apiId string) ([]*apigateway.UpdateRequestValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllDomainNameBasePathMappings(domainName string) ([]*apigateway.BasePathMapping, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllRestApiModels(apiId string) ([]*apigateway.Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *apigatewayRepository) ListAllRestApiGatewayResponses(apiId string) ([]*apigateway.UpdateGatewayResponseOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

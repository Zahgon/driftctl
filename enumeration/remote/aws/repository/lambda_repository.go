package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type LambdaRepository interface {
	ListAllLambdaFunctions() ([]*lambda.FunctionConfiguration, error)
	ListAllLambdaEventSourceMappings() ([]*lambda.EventSourceMappingConfiguration, error)
}

type lambdaRepository struct {
	client lambdaiface.LambdaAPI
	cache  cache.Cache
}

func NewLambdaRepository(session *session.Session, c cache.Cache) *lambdaRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *lambdaRepository) ListAllLambdaFunctions() ([]*lambda.FunctionConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *lambdaRepository) ListAllLambdaEventSourceMappings() ([]*lambda.EventSourceMappingConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

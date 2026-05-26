package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// AwsDefaultApiGatewayAccount is a middleware that ignores the default API Gateway account resource in the current region.
type AwsDefaultApiGatewayAccount struct{}

func NewAwsDefaultApiGatewayAccount() AwsDefaultApiGatewayAccount {
	_ = "STUB: not implemented"
	return *new(AwsDefaultApiGatewayAccount)
}

func (m AwsDefaultApiGatewayAccount) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than API gateway account

// Check if account is managed by IaC

// Include resource if it's managed in IaC

// Else, resource is not added to newRemoteResources slice, so it will be ignored

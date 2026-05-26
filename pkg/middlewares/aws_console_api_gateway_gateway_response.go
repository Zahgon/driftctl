package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Each API Gateway rest API has by design all the gateway responses available to edit in the console
// which result in useless noises (e.g. lots of unmanaged resources) by driftctl.
// This middleware ignores all console responses if not managed by IAC.
type AwsConsoleApiGatewayGatewayResponse struct{}

func NewAwsConsoleApiGatewayGatewayResponse() AwsConsoleApiGatewayGatewayResponse {
	_ = "STUB: not implemented"
	return *new(AwsConsoleApiGatewayGatewayResponse)
}

func (m AwsConsoleApiGatewayGatewayResponse) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than gateway responses

// Check if gateway response is managed by IaC

// Include resource if it's managed by IaC

// Else, resource is not added to newRemoteResources slice so it will be ignored

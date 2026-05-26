package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type AwsInstanceEIP struct{}

func (a AwsInstanceEIP) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than aws_instance

func (a AwsInstanceEIP) hasEIP(instance *resource.Resource, resources *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

func (a AwsInstanceEIP) ignorePublicIpAndDns(instance *resource.Resource, resourcesSet ...*[]*resource.Resource) {
	_ = "STUB: not implemented"
	return
}

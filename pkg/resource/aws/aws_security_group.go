package aws

import (
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

const AwsSecurityGroupResourceType = "aws_security_group"

func initAwsSecurityGroupMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	_ = "STUB: not implemented"
	return
}

// TODO We need to find a way to warn users that some rules in their states could be unmanaged

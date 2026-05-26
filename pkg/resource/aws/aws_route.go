package aws

import (
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

const AwsRouteResourceType = "aws_route"

func initAwsRouteMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	_ = "STUB: not implemented"
	return
}

func CalculateRouteID(tableId, CidrBlock, Ipv6CidrBlock, PrefixListId *string) string {
	_ = "STUB: not implemented"
	return ""
}

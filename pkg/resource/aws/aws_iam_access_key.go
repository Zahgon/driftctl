package aws

import (
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

const AwsIamAccessKeyResourceType = "aws_iam_access_key"

func initAwsIAMAccessKeyMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	_ = "STUB: not implemented"
	return
}

// As we can't read secrets from aws API once access_key created we need to set
// fields retrieved from state to nil to avoid drift
// We can't detect drift if we cannot retrieve latest value from aws API for fields like secrets, passwords etc ...

package aws

import (
	dctlresource "github.com/snyk/driftctl/pkg/resource"

	"github.com/snyk/driftctl/enumeration/resource"
)

const AwsSecurityGroupRuleResourceType = "aws_security_group_rule"

func initAwsSecurityGroupRuleMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	_ = "STUB: not implemented"
	return
}

// On first run, this field is set to null in state file and to "" after one refresh or apply
// This ensure that if we find a nil value we dont drift

// If protocol is all (e.g. -1), tcp, udp, icmp or icmpv6 then we leave the resource untouched
// Else we delete the FromPort/ToPort and recreate the rule's id

func join(elems []interface{}, sep string) string { _ = "STUB: not implemented"; return "" }

func CreateSecurityGroupRuleIdHash(attrs *resource.Attributes) string {
	_ = "STUB: not implemented"
	return ""
}

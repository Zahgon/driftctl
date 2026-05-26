package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"

	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	sgRuleTypeIngress = "ingress"
	sgRuleTypeEgress  = "egress"
)

type VPCSecurityGroupRuleEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

type securityGroupRule struct {
	Type                  string
	SecurityGroupId       string
	Protocol              string
	FromPort              float64
	ToPort                float64
	Self                  bool
	SourceSecurityGroupId string
	CidrBlocks            []string
	Ipv6CidrBlocks        []string
	PrefixListIds         []string
}

func (s *securityGroupRule) getId() string { _ = "STUB: not implemented"; return "" }

func (s *securityGroupRule) getAttrs() resource.Attributes {
	_ = "STUB: not implemented"
	return *new(resource.Attributes)
}

func toInterfaceSlice(val []string) []interface{} { _ = "STUB: not implemented"; return nil }

func NewVPCSecurityGroupRuleEnumerator(repository repository.EC2Repository, factory resource.ResourceFactory) *VPCSecurityGroupRuleEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *VPCSecurityGroupRuleEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *VPCSecurityGroupRuleEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *VPCSecurityGroupRuleEnumerator) listSecurityGroupsRules(securityGroups []*ec2.SecurityGroup) []securityGroupRule {
	_ = "STUB: not implemented"
	return nil
}

// addSecurityGroupRule will iterate through each "Source" as per Aws definition and create a
// rule with custom attributes
func (e *VPCSecurityGroupRuleEnumerator) addSecurityGroupRule(ruleType string, rule *ec2.IpPermission, sg *ec2.SecurityGroup) []securityGroupRule {
	_ = "STUB: not implemented"
	return nil
}

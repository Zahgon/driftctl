package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type EC2Repository interface {
	ListAllImages() ([]*ec2.Image, error)
	ListAllSnapshots() ([]*ec2.Snapshot, error)
	ListAllVolumes() ([]*ec2.Volume, error)
	ListAllAddresses() ([]*ec2.Address, error)
	ListAllAddressesAssociation() ([]*ec2.Address, error)
	ListAllInstances() ([]*ec2.Instance, error)
	ListAllKeyPairs() ([]*ec2.KeyPairInfo, error)
	ListAllInternetGateways() ([]*ec2.InternetGateway, error)
	ListAllSubnets() ([]*ec2.Subnet, []*ec2.Subnet, error)
	ListAllNatGateways() ([]*ec2.NatGateway, error)
	ListAllRouteTables() ([]*ec2.RouteTable, error)
	ListAllVPCs() ([]*ec2.Vpc, []*ec2.Vpc, error)
	ListAllSecurityGroups() ([]*ec2.SecurityGroup, []*ec2.SecurityGroup, error)
	ListAllNetworkACLs() ([]*ec2.NetworkAcl, error)
	DescribeLaunchTemplates() ([]*ec2.LaunchTemplate, error)
	IsEbsEncryptionEnabledByDefault() (bool, error)
}

type ec2Repository struct {
	client ec2iface.EC2API
	cache  cache.Cache
}

func NewEC2Repository(session *session.Session, c cache.Cache) *ec2Repository {
	_ = "STUB: not implemented"
	return nil
}

func (r *ec2Repository) ListAllImages() ([]*ec2.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllSnapshots() ([]*ec2.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllVolumes() ([]*ec2.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllAddresses() ([]*ec2.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllAddressesAssociation() ([]*ec2.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllInstances() ([]*ec2.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignore terminated state from enumeration since terminated means that instance
// has been removed

func (r *ec2Repository) ListAllKeyPairs() ([]*ec2.KeyPairInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllInternetGateways() ([]*ec2.InternetGateway, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllSubnets() ([]*ec2.Subnet, []*ec2.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *ec2Repository) ListAllNatGateways() ([]*ec2.NatGateway, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllRouteTables() ([]*ec2.RouteTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) ListAllVPCs() ([]*ec2.Vpc, []*ec2.Vpc, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *ec2Repository) ListAllSecurityGroups() ([]*ec2.SecurityGroup, []*ec2.SecurityGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *ec2Repository) ListAllNetworkACLs() ([]*ec2.NetworkAcl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) DescribeLaunchTemplates() ([]*ec2.LaunchTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ec2Repository) IsEbsEncryptionEnabledByDefault() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

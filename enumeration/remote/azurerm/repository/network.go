package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/common"
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

type NetworkRepository interface {
	ListAllVirtualNetworks() ([]*armnetwork.VirtualNetwork, error)
	ListAllRouteTables() ([]*armnetwork.RouteTable, error)
	ListAllSubnets(virtualNetwork *armnetwork.VirtualNetwork) ([]*armnetwork.Subnet, error)
	ListAllFirewalls() ([]*armnetwork.AzureFirewall, error)
	ListAllPublicIPAddresses() ([]*armnetwork.PublicIPAddress, error)
	ListAllSecurityGroups() ([]*armnetwork.NetworkSecurityGroup, error)
	ListAllLoadBalancers() ([]*armnetwork.LoadBalancer, error)
	ListLoadBalancerRules(*armnetwork.LoadBalancer) ([]*armnetwork.LoadBalancingRule, error)
}

type publicIPAddressesClient interface {
	ListAll(options *armnetwork.PublicIPAddressesListAllOptions) publicIPAddressesListAllPager
}

type publicIPAddressesListAllPager interface {
	pager
	PageResponse() armnetwork.PublicIPAddressesListAllResponse
}

type publicIPAddressesClientImpl struct {
	client *armnetwork.PublicIPAddressesClient
}

func (p publicIPAddressesClientImpl) ListAll(options *armnetwork.PublicIPAddressesListAllOptions) publicIPAddressesListAllPager {
	_ = "STUB: not implemented"
	return *new(publicIPAddressesListAllPager)
}

type firewallsListAllPager interface {
	pager
	PageResponse() armnetwork.AzureFirewallsListAllResponse
}

type firewallsClient interface {
	ListAll(options *armnetwork.AzureFirewallsListAllOptions) firewallsListAllPager
}

type firewallsClientImpl struct {
	client *armnetwork.AzureFirewallsClient
}

func (s firewallsClientImpl) ListAll(options *armnetwork.AzureFirewallsListAllOptions) firewallsListAllPager {
	_ = "STUB: not implemented"
	return *new(firewallsListAllPager)
}

type subnetsListPager interface {
	pager
	PageResponse() armnetwork.SubnetsListResponse
}

type subnetsClient interface {
	List(resourceGroupName, virtualNetworkName string, options *armnetwork.SubnetsListOptions) subnetsListPager
}

type subnetsClientImpl struct {
	client *armnetwork.SubnetsClient
}

func (s subnetsClientImpl) List(resourceGroupName, virtualNetworkName string, options *armnetwork.SubnetsListOptions) subnetsListPager {
	_ = "STUB: not implemented"
	return *new(subnetsListPager)
}

type virtualNetworksClient interface {
	ListAll(options *armnetwork.VirtualNetworksListAllOptions) virtualNetworksListAllPager
}

type virtualNetworksListAllPager interface {
	pager
	PageResponse() armnetwork.VirtualNetworksListAllResponse
}

type virtualNetworksClientImpl struct {
	client *armnetwork.VirtualNetworksClient
}

func (c virtualNetworksClientImpl) ListAll(options *armnetwork.VirtualNetworksListAllOptions) virtualNetworksListAllPager {
	_ = "STUB: not implemented"
	return *new(virtualNetworksListAllPager)
}

type routeTablesClient interface {
	ListAll(options *armnetwork.RouteTablesListAllOptions) routeTablesListAllPager
}

type routeTablesListAllPager interface {
	pager
	PageResponse() armnetwork.RouteTablesListAllResponse
}

type routeTablesClientImpl struct {
	client *armnetwork.RouteTablesClient
}

func (c routeTablesClientImpl) ListAll(options *armnetwork.RouteTablesListAllOptions) routeTablesListAllPager {
	_ = "STUB: not implemented"
	return *new(routeTablesListAllPager)
}

type networkSecurityGroupsListAllPager interface {
	pager
	PageResponse() armnetwork.NetworkSecurityGroupsListAllResponse
}

type networkSecurityGroupsClient interface {
	ListAll(options *armnetwork.NetworkSecurityGroupsListAllOptions) networkSecurityGroupsListAllPager
}

type networkSecurityGroupsClientImpl struct {
	client *armnetwork.NetworkSecurityGroupsClient
}

func (s networkSecurityGroupsClientImpl) ListAll(options *armnetwork.NetworkSecurityGroupsListAllOptions) networkSecurityGroupsListAllPager {
	_ = "STUB: not implemented"
	return *new(networkSecurityGroupsListAllPager)
}

type loadBalancersListAllPager interface {
	pager
	PageResponse() armnetwork.LoadBalancersListAllResponse
}

type loadBalancersClient interface {
	ListAll(options *armnetwork.LoadBalancersListAllOptions) loadBalancersListAllPager
}

type loadBalancersClientImpl struct {
	client *armnetwork.LoadBalancersClient
}

func (s loadBalancersClientImpl) ListAll(options *armnetwork.LoadBalancersListAllOptions) loadBalancersListAllPager {
	_ = "STUB: not implemented"
	return *new(loadBalancersListAllPager)
}

type loadBalancerRulesListAllPager interface {
	pager
	PageResponse() armnetwork.LoadBalancerLoadBalancingRulesListResponse
}

type loadBalancerRulesClient interface {
	List(string, string, *armnetwork.LoadBalancerLoadBalancingRulesListOptions) loadBalancerRulesListAllPager
}

type loadBalancerRulesClientImpl struct {
	client *armnetwork.LoadBalancerLoadBalancingRulesClient
}

func (s loadBalancerRulesClientImpl) List(resourceGroupName string, loadBalancerName string, options *armnetwork.LoadBalancerLoadBalancingRulesListOptions) loadBalancerRulesListAllPager {
	_ = "STUB: not implemented"
	return *new(loadBalancerRulesListAllPager)
}

type networkRepository struct {
	virtualNetworksClient       virtualNetworksClient
	routeTableClient            routeTablesClient
	subnetsClient               subnetsClient
	firewallsClient             firewallsClient
	publicIPAddressesClient     publicIPAddressesClient
	networkSecurityGroupsClient networkSecurityGroupsClient
	loadBalancersClient         loadBalancersClient
	loadBalancerRulesClient     loadBalancerRulesClient
	cache                       cache.Cache
}

func NewNetworkRepository(cred azcore.TokenCredential, options *arm.ClientOptions, config common.AzureProviderConfig, cache cache.Cache) *networkRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *networkRepository) ListAllVirtualNetworks() ([]*armnetwork.VirtualNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *networkRepository) ListAllRouteTables() ([]*armnetwork.RouteTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *networkRepository) ListAllSubnets(virtualNetwork *armnetwork.VirtualNetwork) ([]*armnetwork.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *networkRepository) ListAllFirewalls() ([]*armnetwork.AzureFirewall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *networkRepository) ListAllPublicIPAddresses() ([]*armnetwork.PublicIPAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *networkRepository) ListAllSecurityGroups() ([]*armnetwork.NetworkSecurityGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *networkRepository) ListAllLoadBalancers() ([]*armnetwork.LoadBalancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *networkRepository) ListLoadBalancerRules(loadBalancer *armnetwork.LoadBalancer) ([]*armnetwork.LoadBalancingRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

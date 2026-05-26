package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/common"
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

type ResourcesRepository interface {
	ListAllResourceGroups() ([]*armresources.ResourceGroup, error)
}

type resourcesListPager interface {
	pager
	PageResponse() armresources.ResourceGroupsListResponse
}

type resourcesClient interface {
	List(options *armresources.ResourceGroupsListOptions) resourcesListPager
}

type resourcesClientImpl struct {
	client *armresources.ResourceGroupsClient
}

func (c resourcesClientImpl) List(options *armresources.ResourceGroupsListOptions) resourcesListPager {
	_ = "STUB: not implemented"
	return *new(resourcesListPager)
}

type resourcesRepository struct {
	client resourcesClient
	cache  cache.Cache
}

func NewResourcesRepository(cred azcore.TokenCredential, options *arm.ClientOptions, config common.AzureProviderConfig, cache cache.Cache) *resourcesRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *resourcesRepository) ListAllResourceGroups() ([]*armresources.ResourceGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

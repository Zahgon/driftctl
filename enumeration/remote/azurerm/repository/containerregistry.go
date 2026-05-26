package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/common"
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

type ContainerRegistryRepository interface {
	ListAllContainerRegistries() ([]*armcontainerregistry.Registry, error)
}

type registryClient interface {
	List(options *armcontainerregistry.RegistriesListOptions) registryListAllPager
}

type registryListAllPager interface {
	pager
	PageResponse() armcontainerregistry.RegistriesListResponse
}

type registryClientImpl struct {
	client *armcontainerregistry.RegistriesClient
}

func (c registryClientImpl) List(options *armcontainerregistry.RegistriesListOptions) registryListAllPager {
	_ = "STUB: not implemented"
	return *new(registryListAllPager)
}

type containerRegistryRepository struct {
	registryClient registryClient
	cache          cache.Cache
}

func NewContainerRegistryRepository(cred azcore.TokenCredential, options *arm.ClientOptions, config common.AzureProviderConfig, cache cache.Cache) *containerRegistryRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *containerRegistryRepository) ListAllContainerRegistries() ([]*armcontainerregistry.Registry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/common"
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

type StorageRespository interface {
	ListAllStorageAccount() ([]*armstorage.StorageAccount, error)
	ListAllStorageContainer(account *armstorage.StorageAccount) ([]string, error)
}

type blobContainerListPager interface {
	pager
	PageResponse() armstorage.BlobContainersListResponse
}

// Interfaces are only used to create mock on Azure SDK
type blobContainerClient interface {
	List(resourceGroupName string, accountName string, options *armstorage.BlobContainersListOptions) blobContainerListPager
}

type blobContainerClientImpl struct {
	client *armstorage.BlobContainersClient
}

func (c blobContainerClientImpl) List(resourceGroupName string, accountName string, options *armstorage.BlobContainersListOptions) blobContainerListPager {
	_ = "STUB: not implemented"
	return *new(blobContainerListPager)
}

type storageAccountListPager interface {
	pager
	PageResponse() armstorage.StorageAccountsListResponse
}

type storageAccountClient interface {
	List(options *armstorage.StorageAccountsListOptions) storageAccountListPager
}

type storageAccountClientImpl struct {
	client *armstorage.StorageAccountsClient
}

func (c storageAccountClientImpl) List(options *armstorage.StorageAccountsListOptions) storageAccountListPager {
	_ = "STUB: not implemented"
	return *new(storageAccountListPager)
}

type storageRepository struct {
	storageAccountsClient storageAccountClient
	blobContainerClient   blobContainerClient
	cache                 cache.Cache
}

func NewStorageRepository(cred azcore.TokenCredential, options *arm.ClientOptions, config common.AzureProviderConfig, cache cache.Cache) *storageRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *storageRepository) ListAllStorageAccount() ([]*armstorage.StorageAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *storageRepository) ListAllStorageContainer(account *armstorage.StorageAccount) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func shouldIgnoreStorageContainerError(err error) bool { _ = "STUB: not implemented"; return false }

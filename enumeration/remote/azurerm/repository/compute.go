package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/common"
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

type ComputeRepository interface {
	ListAllImages() ([]*armcompute.Image, error)
	ListAllSSHPublicKeys() ([]*armcompute.SSHPublicKeyResource, error)
}

type imagesListPager interface {
	pager
	PageResponse() armcompute.ImagesListResponse
}

type imagesClient interface {
	List(options *armcompute.ImagesListOptions) imagesListPager
}

type imagesClientImpl struct {
	client *armcompute.ImagesClient
}

func (c imagesClientImpl) List(options *armcompute.ImagesListOptions) imagesListPager {
	_ = "STUB: not implemented"
	return *new(imagesListPager)
}

type sshPublicKeyListPager interface {
	pager
	PageResponse() armcompute.SSHPublicKeysListBySubscriptionResponse
}

type sshPublicKeyClient interface {
	ListBySubscription(options *armcompute.SSHPublicKeysListBySubscriptionOptions) sshPublicKeyListPager
}

type sshPublicKeyClientImpl struct {
	client *armcompute.SSHPublicKeysClient
}

func (c sshPublicKeyClientImpl) ListBySubscription(options *armcompute.SSHPublicKeysListBySubscriptionOptions) sshPublicKeyListPager {
	_ = "STUB: not implemented"
	return *new(sshPublicKeyListPager)
}

type computeRepository struct {
	imagesClient       imagesClient
	sshPublicKeyClient sshPublicKeyClient
	cache              cache.Cache
}

func NewComputeRepository(cred azcore.TokenCredential, options *arm.ClientOptions, config common.AzureProviderConfig, cache cache.Cache) *computeRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *computeRepository) ListAllImages() ([]*armcompute.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *computeRepository) ListAllSSHPublicKeys() ([]*armcompute.SSHPublicKeyResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

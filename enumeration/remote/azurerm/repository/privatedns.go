package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/common"
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

type PrivateDNSRepository interface {
	ListAllPrivateZones() ([]*armprivatedns.PrivateZone, error)
	ListAllARecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error)
	ListAllAAAARecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error)
	ListAllCNAMERecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error)
	ListAllPTRRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error)
	ListAllMXRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error)
	ListAllSRVRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error)
	ListAllTXTRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error)
}

type privateDNSZoneListPager interface {
	pager
	PageResponse() armprivatedns.PrivateZonesListResponse
}

type privateDNSRecordSetListPager interface {
	pager
	PageResponse() armprivatedns.RecordSetsListResponse
}

type privateRecordSetClient interface {
	List(resourceGroupName string, privateZoneName string, options *armprivatedns.RecordSetsListOptions) privateDNSRecordSetListPager
}

type privateRecordSetClientImpl struct {
	client *armprivatedns.RecordSetsClient
}

func (c *privateRecordSetClientImpl) List(resourceGroupName string, privateZoneName string, options *armprivatedns.RecordSetsListOptions) privateDNSRecordSetListPager {
	_ = "STUB: not implemented"
	return *new(privateDNSRecordSetListPager)
}

type privateZonesClient interface {
	List(options *armprivatedns.PrivateZonesListOptions) privateDNSZoneListPager
}

type privateZonesClientImpl struct {
	client *armprivatedns.PrivateZonesClient
}

func (c *privateZonesClientImpl) List(options *armprivatedns.PrivateZonesListOptions) privateDNSZoneListPager {
	_ = "STUB: not implemented"
	return *new(privateDNSZoneListPager)
}

type privateDNSRepository struct {
	zoneClient   privateZonesClient
	recordClient privateRecordSetClient
	cache        cache.Cache
}

func NewPrivateDNSRepository(cred azcore.TokenCredential, options *arm.ClientOptions, config common.AzureProviderConfig, cache cache.Cache) *privateDNSRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *privateDNSRepository) listAllRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *privateDNSRepository) ListAllARecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *privateDNSRepository) ListAllAAAARecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *privateDNSRepository) ListAllPTRRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *privateDNSRepository) ListAllCNAMERecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *privateDNSRepository) ListAllMXRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *privateDNSRepository) ListAllSRVRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *privateDNSRepository) ListAllTXTRecords(zone *armprivatedns.PrivateZone) ([]*armprivatedns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *privateDNSRepository) ListAllPrivateZones() ([]*armprivatedns.PrivateZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

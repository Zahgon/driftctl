package repository

import (
	asset "cloud.google.com/go/asset/apiv1"
	assetpb "cloud.google.com/go/asset/apiv1/assetpb"
	"github.com/snyk/driftctl/enumeration/remote/cache"
	"github.com/snyk/driftctl/enumeration/remote/google/config"
)

// https://cloud.google.com/asset-inventory/docs/supported-asset-types#supported_resource_types
const (
	storageBucketAssetType               = "storage.googleapis.com/Bucket"
	computeFirewallAssetType             = "compute.googleapis.com/Firewall"
	computeRouterAssetType               = "compute.googleapis.com/Router"
	computeInstanceAssetType             = "compute.googleapis.com/Instance"
	computeNetworkAssetType              = "compute.googleapis.com/Network"
	computeSubnetworkAssetType           = "compute.googleapis.com/Subnetwork"
	computeDiskAssetType                 = "compute.googleapis.com/Disk"
	computeImageAssetType                = "compute.googleapis.com/Image"
	dnsManagedZoneAssetType              = "dns.googleapis.com/ManagedZone"
	computeInstanceGroupAssetType        = "compute.googleapis.com/InstanceGroup"
	bigqueryDatasetAssetType             = "bigquery.googleapis.com/Dataset"
	bigqueryTableAssetType               = "bigquery.googleapis.com/Table"
	computeAddressAssetType              = "compute.googleapis.com/Address"
	computeGlobalAddressAssetType        = "compute.googleapis.com/GlobalAddress"
	cloudFunctionsFunction               = "cloudfunctions.googleapis.com/CloudFunction"
	bigtableInstanceAssetType            = "bigtableadmin.googleapis.com/Instance"
	bigtableTableAssetType               = "bigtableadmin.googleapis.com/Table"
	sqlDatabaseInstanceAssetType         = "sqladmin.googleapis.com/Instance"
	healthCheckAssetType                 = "compute.googleapis.com/HealthCheck"
	cloudRunServiceAssetType             = "run.googleapis.com/Service"
	nodeGroupAssetType                   = "compute.googleapis.com/NodeGroup"
	computeForwardingRuleAssetType       = "compute.googleapis.com/ForwardingRule"
	instanceGroupManagerAssetType        = "compute.googleapis.com/InstanceGroupManager"
	computeGlobalForwardingRuleAssetType = "compute.googleapis.com/GlobalForwardingRule"
	computeSslCertificateAssetType       = "compute.googleapis.com/SslCertificate"
)

type AssetRepository interface {
	SearchAllBuckets() ([]*assetpb.ResourceSearchResult, error)
	SearchAllFirewalls() ([]*assetpb.ResourceSearchResult, error)
	SearchAllRouters() ([]*assetpb.ResourceSearchResult, error)
	SearchAllInstances() ([]*assetpb.ResourceSearchResult, error)
	SearchAllNetworks() ([]*assetpb.ResourceSearchResult, error)
	SearchAllDisks() ([]*assetpb.ResourceSearchResult, error)
	SearchAllImages() ([]*assetpb.ResourceSearchResult, error)
	SearchAllDNSManagedZones() ([]*assetpb.ResourceSearchResult, error)
	SearchAllInstanceGroups() ([]*assetpb.ResourceSearchResult, error)
	SearchAllDatasets() ([]*assetpb.ResourceSearchResult, error)
	SearchAllTables() ([]*assetpb.ResourceSearchResult, error)
	SearchAllAddresses() ([]*assetpb.ResourceSearchResult, error)
	SearchAllGlobalAddresses() ([]*assetpb.Asset, error)
	SearchAllFunctions() ([]*assetpb.Asset, error)
	SearchAllSubnetworks() ([]*assetpb.ResourceSearchResult, error)
	SearchAllBigtableInstances() ([]*assetpb.Asset, error)
	SearchAllBigtableTables() ([]*assetpb.Asset, error)
	SearchAllSQLDatabaseInstances() ([]*assetpb.Asset, error)
	SearchAllHealthChecks() ([]*assetpb.ResourceSearchResult, error)
	SearchAllCloudRunServices() ([]*assetpb.ResourceSearchResult, error)
	SearchAllNodeGroups() ([]*assetpb.Asset, error)
	SearchAllForwardingRules() ([]*assetpb.Asset, error)
	SearchAllInstanceGroupManagers() ([]*assetpb.Asset, error)
	SearchAllGlobalForwardingRules() ([]*assetpb.Asset, error)
	SearchAllSslCertificates() ([]*assetpb.Asset, error)
}

type assetRepository struct {
	client *asset.Client
	config config.GCPTerraformConfig
	cache  cache.Cache
}

func NewAssetRepository(client *asset.Client, config config.GCPTerraformConfig, c cache.Cache) *assetRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s assetRepository) listAllResources(ty string) ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) searchAllResources(ty string) ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllBuckets() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllFirewalls() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllRouters() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllInstances() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllNetworks() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllDNSManagedZones() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllInstanceGroups() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllDatasets() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllTables() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllAddresses() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllGlobalAddresses() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllFunctions() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllSubnetworks() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllDisks() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllImages() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllBigtableInstances() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllBigtableTables() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllSQLDatabaseInstances() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllHealthChecks() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllCloudRunServices() ([]*assetpb.ResourceSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllNodeGroups() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllForwardingRules() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllInstanceGroupManagers() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllGlobalForwardingRules() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s assetRepository) SearchAllSslCertificates() ([]*assetpb.Asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

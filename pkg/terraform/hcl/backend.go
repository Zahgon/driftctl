package hcl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/snyk/driftctl/pkg/iac/config"
)

type BackendBlock struct {
	Name               string   `hcl:"name,label"`
	Path               string   `hcl:"path,optional"`
	WorkspaceDir       string   `hcl:"workspace_dir,optional"`
	Bucket             string   `hcl:"bucket,optional"`
	Key                string   `hcl:"key,optional"`
	Region             string   `hcl:"region,optional"`
	Prefix             string   `hcl:"prefix,optional"`
	ContainerName      string   `hcl:"container_name,optional"`
	WorkspaceKeyPrefix string   `hcl:"workspace_key_prefix,optional"`
	Remain             hcl.Body `hcl:",remain"`
}

func (b BackendBlock) SupplierConfig(workspace string) *config.SupplierConfig {
	_ = "STUB: not implemented"
	return nil
}

func (b BackendBlock) parseLocalBackend() *config.SupplierConfig {
	_ = "STUB: not implemented"
	return nil
}

func (b BackendBlock) parseS3Backend(ws string) *config.SupplierConfig {
	_ = "STUB: not implemented"
	return nil
}

func (b BackendBlock) parseGCSBackend(ws string) *config.SupplierConfig {
	_ = "STUB: not implemented"
	return nil
}

func (b BackendBlock) parseAzurermBackend(ws string) *config.SupplierConfig {
	_ = "STUB: not implemented"
	return nil
}

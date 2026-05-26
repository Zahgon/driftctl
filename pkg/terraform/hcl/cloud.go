package hcl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/snyk/driftctl/pkg/iac/config"
)

type CloudWorkspacesBlock struct {
	Name   string   `hcl:"name,optional"`
	Tags   []string `hcl:"tags,optional"`
	Remain hcl.Body `hcl:",remain"`
}

type CloudBlock struct {
	Organization string               `hcl:"organization"`
	Workspaces   CloudWorkspacesBlock `hcl:"workspaces,block"`
	Remain       hcl.Body             `hcl:",remain"`
}

func (c CloudBlock) SupplierConfig(workspace string) *config.SupplierConfig {
	_ = "STUB: not implemented"
	// If a workspace is specified in HCL, use it rather than the current environment
	return nil
}

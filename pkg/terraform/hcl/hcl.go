package hcl

import (
	"github.com/hashicorp/hcl/v2"
)

const DefaultStateName = "default"

type MainBodyBlock struct {
	Terraform TerraformBlock `hcl:"terraform,block"`
	Remain    hcl.Body       `hcl:",remain"`
}

type TerraformBlock struct {
	Backend *BackendBlock `hcl:"backend,block"`
	Cloud   *CloudBlock   `hcl:"cloud,block"`
	Remain  hcl.Body      `hcl:",remain"`
}

func ParseTerraformFromHCL(filename string) (*TerraformBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetCurrentWorkspaceName(cwd string) string { _ = "STUB: not implemented"; return "" }

// See https://github.com/hashicorp/terraform/blob/main/internal/backend/backend.go#L33

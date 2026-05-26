package terraform

import (
	"github.com/hashicorp/terraform/providers"
	"github.com/snyk/driftctl/enumeration/terraform"
	"github.com/zclconf/go-cty/cty"
)

type FakeTerraformProvider struct {
	realProvider terraform.TerraformProvider
	shouldUpdate bool
	response     string
}

func NewFakeTerraformProvider(realProvider terraform.TerraformProvider) *FakeTerraformProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *FakeTerraformProvider) ShouldUpdate() { _ = "STUB: not implemented"; return }

func (p *FakeTerraformProvider) Schema() map[string]providers.Schema {
	_ = "STUB: not implemented"
	return nil
}

func (p *FakeTerraformProvider) WithResponse(response string) *FakeTerraformProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *FakeTerraformProvider) ReadResource(args terraform.ReadResourceArgs) (*cty.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *FakeTerraformProvider) readSchema() map[string]providers.Schema {
	_ = "STUB: not implemented"
	return nil
}

func (p *FakeTerraformProvider) writeResource(args terraform.ReadResourceArgs, readResource *cty.Value, err error) {
	_ = "STUB: not implemented"
	return
}

func (p *FakeTerraformProvider) readResource(args terraform.ReadResourceArgs) (*cty.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *FakeTerraformProvider) getFileName(args terraform.ReadResourceArgs) string {
	_ = "STUB: not implemented"
	return ""
}

// ext4 and many other filesystems has a maximum filename length of 255 bytes
// See https://en.wikipedia.org/wiki/Comparison_of_file_systems#Limits
// Solution: we create a SHA1 hash of the filename so the length stay constant
// We should do that no matter the length, but it requires to regenerate every single file
// TODO: Use SHA1 filenames for all resource golden files

func (p *FakeTerraformProvider) Cleanup() { _ = "STUB: not implemented"; return }

func (p *FakeTerraformProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *FakeTerraformProvider) Version() string { _ = "STUB: not implemented"; return "" }

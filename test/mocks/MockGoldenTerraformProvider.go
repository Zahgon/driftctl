package mocks

import (
	terraform2 "github.com/snyk/driftctl/enumeration/terraform"

	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
)

type MockedGoldenTFProvider struct {
	name            string
	providerName    string
	providerVersion string
	realProvider    terraform2.TerraformProvider
	update          bool
}

func NewMockedGoldenTFProvider(name, providerName, providerVersion string, realProvider terraform2.TerraformProvider, update bool) *MockedGoldenTFProvider {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockedGoldenTFProvider) Schema() map[string]providers.Schema {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockedGoldenTFProvider) ReadResource(args terraform2.ReadResourceArgs) (*cty.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockedGoldenTFProvider) writeSchema(schema map[string]providers.Schema) {
	_ = "STUB: not implemented"
	return
}

func (m *MockedGoldenTFProvider) getSchemaPath() string { _ = "STUB: not implemented"; return "" }

func (m *MockedGoldenTFProvider) readSchema() map[string]providers.Schema {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockedGoldenTFProvider) writeReadResource(args terraform2.ReadResourceArgs, readResource *cty.Value, err error) {
	_ = "STUB: not implemented"
	return
}

func (m *MockedGoldenTFProvider) readReadResource(args terraform2.ReadResourceArgs) (*cty.Value, error) {
	_ = "STUB: not implemented"
	return nil,

		// TODO I'm putting this here for compatibility reason...
		nil
}

type ReadResource struct {
	Value *cty.Value
	Err   error
}

func (m *ReadResource) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (m *ReadResource) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func getFileName(args terraform2.ReadResourceArgs) string { _ = "STUB: not implemented"; return "" }

func getFileNameSuffix(args terraform2.ReadResourceArgs) string {
	_ = "STUB: not implemented"
	return ""
}

func (p MockedGoldenTFProvider) Cleanup() { _ = "STUB: not implemented"; return }

func (p *MockedGoldenTFProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *MockedGoldenTFProvider) Version() string { _ = "STUB: not implemented"; return "" }

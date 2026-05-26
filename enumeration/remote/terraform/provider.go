package terraform

import (
	"sync"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/providers"
	progress2 "github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/parallel"
	tf "github.com/snyk/driftctl/enumeration/terraform"
	"github.com/zclconf/go-cty/cty"
)

const EXIT_ERROR = 3

// "alias" in these struct are a way to namespace gRPC clients.
// For example, if we need to read S3 bucket from multiple AWS region,
// we'll have an alias per region, and the alias IS the region itself.
// So we can query resources using a specific custom provider configuration
type TerraformProviderConfig struct {
	Name              string
	DefaultAlias      string
	GetProviderConfig func(alias string) interface{}
}

type TerraformProvider struct {
	lock              sync.Mutex
	providerInstaller *tf.ProviderInstaller
	grpcProviders     map[string]*plugin.GRPCProvider
	schemas           map[string]providers.Schema
	Config            TerraformProviderConfig
	runner            *parallel.ParallelRunner
	progress          progress2.ProgressCounter
}

func NewTerraformProvider(installer *tf.ProviderInstaller, config TerraformProviderConfig, progress progress2.ProgressCounter) (*TerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *TerraformProvider) Init() error { _ = "STUB: not implemented"; return nil }

func (p *TerraformProvider) Schema() map[string]providers.Schema {
	_ = "STUB: not implemented"
	return nil
}

func (p *TerraformProvider) Runner() *parallel.ParallelRunner {
	_ = "STUB: not implemented"
	return nil
}

func (p *TerraformProvider) configure(alias string) error { _ = "STUB: not implemented"; return nil }

// This value is optional. It'll be overridden by the provider config.

func (p *TerraformProvider) ReadResource(args tf.ReadResourceArgs) (*cty.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// call to the provider sometimes add and delete field to their attribute this may broke caller so we deep copy attributes

func (p *TerraformProvider) Cleanup() { _ = "STUB: not implemented"; return }

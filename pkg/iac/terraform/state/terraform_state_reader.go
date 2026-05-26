package state

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/terraform"

	"github.com/hashicorp/terraform/states"
	"github.com/snyk/driftctl/pkg/filter"
	"github.com/snyk/driftctl/pkg/output"
	"github.com/zclconf/go-cty/cty"

	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/snyk/driftctl/pkg/iac/config"
	"github.com/snyk/driftctl/pkg/iac/terraform/state/backend"
	"github.com/snyk/driftctl/pkg/iac/terraform/state/enumerator"
)

const TerraformStateReaderSupplier = "tfstate"

type decodedRes struct {
	source resource.Source
	val    cty.Value
}

type TerraformStateReader struct {
	library        *terraform.ProviderLibrary
	config         config.SupplierConfig
	backend        backend.Backend
	enumerator     enumerator.StateEnumerator
	deserializer   *resource.Deserializer
	backendOptions *backend.Options
	progress       output.Progress
	filter         filter.Filter
	alerter        *alerter.Alerter
	sourceCount    uint
}

func (r *TerraformStateReader) initReader() error { _ = "STUB: not implemented"; return nil }

func NewReader(config config.SupplierConfig, library *terraform.ProviderLibrary, backendOpts *backend.Options, progress output.Progress, alerter *alerter.Alerter, deserializer *resource.Deserializer, filter filter.Filter) (*TerraformStateReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TerraformStateReader) retrieve() (map[string][]decodedRes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to do a manual type conversion if we got a path error
// It will allow driftctl to read state generated with a superior version of provider
// than the actually supported one
// by ignoring new fields

func (r *TerraformStateReader) convertInstance(instance *states.ResourceInstanceObjectSrc, ty cty.Type) (*states.ResourceInstanceObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TerraformStateReader) decode(valFromState map[string][]decodedRes) ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TerraformStateReader) Resources() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TerraformStateReader) SourceCount() uint { _ = "STUB: not implemented"; return 0 }

func (r *TerraformStateReader) retrieveForState(path string) ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TerraformStateReader) retrieveMultiplesStates() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// all key failed, throw an error

func read(path string, reader backend.Backend) (*states.State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readState(path string, reader backend.Backend) (*states.State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

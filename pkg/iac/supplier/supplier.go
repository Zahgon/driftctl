package supplier

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/terraform"
	resource2 "github.com/snyk/driftctl/pkg/resource"

	"github.com/snyk/driftctl/pkg/filter"
	"github.com/snyk/driftctl/pkg/iac/config"
	"github.com/snyk/driftctl/pkg/iac/terraform/state/backend"
	"github.com/snyk/driftctl/pkg/output"

	"github.com/snyk/driftctl/pkg/iac/terraform/state"

	"github.com/snyk/driftctl/enumeration/resource"
)

var supportedSuppliers = []string{
	state.TerraformStateReaderSupplier,
}

func IsSupplierSupported(supplierKey string) bool { _ = "STUB: not implemented"; return false }

func GetIACSupplier(configs []config.SupplierConfig,
	library *terraform.ProviderLibrary,
	backendOpts *backend.Options,
	progress output.Progress,
	alerter *alerter.Alerter,
	factory resource.ResourceFactory,
	filter filter.Filter) (resource2.IaCSupplier, error) {
	_ = "STUB: not implemented"
	return *new(resource2.IaCSupplier), nil
}

func GetSupportedSuppliers() []string { _ = "STUB: not implemented"; return nil }

func GetSupportedSchemes() []string { _ = "STUB: not implemented"; return nil }

package pkg

import (
	"github.com/jmespath/go-jmespath"
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/snyk/driftctl/pkg/analyser"
	"github.com/snyk/driftctl/pkg/cmd/scan/output"
	"github.com/snyk/driftctl/pkg/iac/config"
	"github.com/snyk/driftctl/pkg/iac/terraform/state/backend"
	"github.com/snyk/driftctl/pkg/memstore"
	globaloutput "github.com/snyk/driftctl/pkg/output"
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

type FmtOptions struct {
	Output output.OutputConfig
}

type ScanOptions struct {
	Coverage         bool
	Detect           bool
	From             []config.SupplierConfig
	To               string
	Output           []output.OutputConfig
	Filter           *jmespath.JMESPath
	Quiet            bool
	BackendOptions   *backend.Options
	StrictMode       bool
	DisableTelemetry bool
	ProviderVersion  string
	ConfigDir        string
	DriftignorePath  string
	Driftignores     []string
}

type DriftCTL struct {
	remoteSupplier           resource.Supplier
	iacSupplier              dctlresource.IaCSupplier
	alerter                  alerter.AlerterInterface
	analyzer                 *analyser.Analyzer
	resourceFactory          resource.ResourceFactory
	scanProgress             globaloutput.Progress
	iacProgress              globaloutput.Progress
	resourceSchemaRepository dctlresource.SchemaRepositoryInterface
	opts                     *ScanOptions
	store                    memstore.Store
}

func NewDriftCTL(remoteSupplier resource.Supplier,
	iacSupplier dctlresource.IaCSupplier,
	alerter *alerter.Alerter,
	analyzer *analyser.Analyzer,
	resFactory resource.ResourceFactory,
	opts *ScanOptions,
	scanProgress globaloutput.Progress,
	iacProgress globaloutput.Progress,
	resourceSchemaRepository dctlresource.SchemaRepositoryInterface,
	store memstore.Store) *DriftCTL {
	_ = "STUB: not implemented"
	return nil
}

func (d DriftCTL) Run() (*analyser.Analysis, error) { _ = "STUB: not implemented"; return nil, nil }

func (d DriftCTL) Stop() { _ = "STUB: not implemented"; return }

func (d DriftCTL) scan() (remoteResources []*resource.Resource, resourcesFromState []*resource.Resource, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// We do a normalization pass to resources from remote because resource in IaC supplier
// are already created using DriftctlFactory.CreateAbstractResource and thus are already normalized

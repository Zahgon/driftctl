package analyser

import (
	"time"

	"github.com/snyk/driftctl/enumeration/alerter"

	"github.com/snyk/driftctl/enumeration/resource"
)

type Summary struct {
	TotalResources      int  `json:"total_resources"`
	TotalUnmanaged      int  `json:"total_unmanaged"`
	TotalDeleted        int  `json:"total_missing"`
	TotalManaged        int  `json:"total_managed"`
	TotalIaCSourceCount uint `json:"total_iac_source_count"`
}

type Analysis struct {
	unmanaged       []*resource.Resource
	managed         []*resource.Resource
	deleted         []*resource.Resource
	summary         Summary
	alerts          alerter.Alerts
	Duration        time.Duration
	Date            time.Time
	ProviderName    string
	ProviderVersion string
}

type serializableAnalysis struct {
	Summary         Summary                                `json:"summary"`
	Managed         []resource.SerializableResource        `json:"managed"`
	Unmanaged       []resource.SerializableResource        `json:"unmanaged"`
	Deleted         []resource.SerializableResource        `json:"missing"`
	Coverage        int                                    `json:"coverage"`
	Alerts          map[string][]alerter.SerializableAlert `json:"alerts"`
	ProviderName    string                                 `json:"provider_name"`
	ProviderVersion string                                 `json:"provider_version"`
	ScanDuration    uint                                   `json:"scan_duration,omitempty"`
	Date            time.Time                              `json:"date"`
}

type GenDriftIgnoreOptions struct {
	ExcludeUnmanaged bool
	ExcludeDeleted   bool
	ExcludeDrifted   bool
	InputPath        string
	OutputPath       string
}

func NewAnalysis() *Analysis { _ = "STUB: not implemented"; return nil }

func (a Analysis) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Analysis) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// We loose the source type in the serialization process, for now everything is serialized back to a
// TerraformStateSource.
// TODO: Add a discriminator field to be able to serialize back to the right type
// when we'll introduce a new source type

func (a *Analysis) IsSync() bool { _ = "STUB: not implemented"; return false }

func (a *Analysis) AddDeleted(resources ...*resource.Resource) { _ = "STUB: not implemented"; return }

func (a *Analysis) AddUnmanaged(resources ...*resource.Resource) { _ = "STUB: not implemented"; return }

func (a *Analysis) AddManaged(resources ...*resource.Resource) { _ = "STUB: not implemented"; return }

func (a *Analysis) SetAlerts(alerts alerter.Alerts) { _ = "STUB: not implemented"; return }

func (a *Analysis) SetIaCSourceCount(i uint) { _ = "STUB: not implemented"; return }

func (a *Analysis) Coverage() int { _ = "STUB: not implemented"; return 0 }

func (a *Analysis) Managed() []*resource.Resource { _ = "STUB: not implemented"; return nil }

func (a *Analysis) Unmanaged() []*resource.Resource { _ = "STUB: not implemented"; return nil }

func (a *Analysis) Deleted() []*resource.Resource { _ = "STUB: not implemented"; return nil }

func (a *Analysis) Summary() Summary { _ = "STUB: not implemented"; return *new(Summary) }

func (a *Analysis) Alerts() alerter.Alerts { _ = "STUB: not implemented"; return *new(alerter.Alerts) }

func (a *Analysis) SortResources() { _ = "STUB: not implemented"; return }

func (a *Analysis) DriftIgnoreList(opts GenDriftIgnoreOptions) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func escapeKey(line string) string { _ = "STUB: not implemented"; return "" }

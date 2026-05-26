package telemetry

import (
	"github.com/snyk/driftctl/build"
	"github.com/snyk/driftctl/pkg/memstore"
)

type telemetry struct {
	Version        string `json:"version"`
	Os             string `json:"os"`
	Arch           string `json:"arch"`
	TotalResources int    `json:"total_resources"`
	TotalManaged   int    `json:"total_managed"`
	Duration       uint   `json:"duration"`
	ProviderName   string `json:"provider_name"`
	IaCSourceCount uint   `json:"iac_source_count"`
	Client         string `json:"client"`
}

type Telemetry struct {
	build build.BuildInterface
}

func NewTelemetry(build build.BuildInterface) *Telemetry { _ = "STUB: not implemented"; return nil }

func (te Telemetry) SendTelemetry(store memstore.Bucket) { _ = "STUB: not implemented"; return }

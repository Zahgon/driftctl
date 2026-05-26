package build

var env = "dev"

// This flag could be switched to false while building to create a binary without third party network calls
// That mean that following services will be disabled:
// - telemetry
// - version check
var enableUsageReporting = "true"

type BuildInterface interface {
	IsRelease() bool
	IsUsageReportingEnabled() bool
}

type Build struct{}

func (b Build) IsRelease() bool { _ = "STUB: not implemented"; return false }

func (b Build) IsUsageReportingEnabled() bool { _ = "STUB: not implemented"; return false }

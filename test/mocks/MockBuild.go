package mocks

type MockBuild struct {
	Release        bool
	UsageReporting bool
}

func (m MockBuild) IsRelease() bool { _ = "STUB: not implemented"; return false }

func (m MockBuild) IsUsageReportingEnabled() bool { _ = "STUB: not implemented"; return false }

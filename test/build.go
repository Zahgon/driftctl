package test

type Build struct{}

func (b Build) IsRelease() bool { _ = "STUB: not implemented"; return false }

func (b Build) IsUsageReportingEnabled() bool { _ = "STUB: not implemented"; return false }

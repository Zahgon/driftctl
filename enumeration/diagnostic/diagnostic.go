package diagnostic

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/resource"
)

type Diagnostic interface {
	Code() string
	Message() string
	ResourceType() string
	Resource() *resource.Resource
}

type diagnosticImpl struct {
	alert alerter.Alert
}

func (d *diagnosticImpl) Code() string { _ = "STUB: not implemented"; return "" }

func (d *diagnosticImpl) Message() string { _ = "STUB: not implemented"; return "" }

func (d *diagnosticImpl) ResourceType() string { _ = "STUB: not implemented"; return "" }

func (d *diagnosticImpl) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

type Diagnostics []Diagnostic

func FromAlerts(alertMap alerter.Alerts) Diagnostics {
	_ = "STUB: not implemented"
	return *new(Diagnostics)
}

package alerter

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type AlerterInterface interface {
	SendAlert(key string, alert Alert)
}

type Alerter struct {
	alerts   Alerts
	alertsCh chan Alerts
	doneCh   chan bool
}

func NewAlerter() *Alerter { _ = "STUB: not implemented"; return nil }

func (a *Alerter) run() { _ = "STUB: not implemented"; return }

func (a *Alerter) SetAlerts(alerts Alerts) { _ = "STUB: not implemented"; return }

func (a *Alerter) Retrieve() Alerts { _ = "STUB: not implemented"; return *new(Alerts) }

func (a *Alerter) SendAlert(key string, alert Alert) { _ = "STUB: not implemented"; return }

func (a *Alerter) IsResourceIgnored(res *resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Alerter) shouldBeIgnored(alert []Alert) bool { _ = "STUB: not implemented"; return false }

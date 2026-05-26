package alerts

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	remoteerror "github.com/snyk/driftctl/enumeration/remote/error"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ScanningPhase int

const (
	EnumerationPhase ScanningPhase = iota
	DetailsFetchingPhase
)

type RemoteAccessDeniedAlert struct {
	message       string
	provider      string
	scanningPhase ScanningPhase
	resource      *resource.Resource
}

func NewRemoteAccessDeniedAlert(provider string, scanErr *remoteerror.ResourceScanningError, scanningPhase ScanningPhase) *RemoteAccessDeniedAlert {
	_ = "STUB: not implemented"
	return nil
}

func (e *RemoteAccessDeniedAlert) Message() string { _ = "STUB: not implemented"; return "" }

func (e *RemoteAccessDeniedAlert) ShouldIgnoreResource() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *RemoteAccessDeniedAlert) Resource() *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

func (e *RemoteAccessDeniedAlert) GetProviderMessage() string { _ = "STUB: not implemented"; return "" }

func sendRemoteAccessDeniedAlert(provider string, alerter alerter.AlerterInterface, listError *remoteerror.ResourceScanningError, p ScanningPhase) {
	_ = "STUB: not implemented"
	return
}

func SendEnumerationAlert(provider string, alerter alerter.AlerterInterface, listError *remoteerror.ResourceScanningError) {
	_ = "STUB: not implemented"
	return
}

func SendDetailsFetchingAlert(provider string, alerter alerter.AlerterInterface, listError *remoteerror.ResourceScanningError) {
	_ = "STUB: not implemented"
	return
}

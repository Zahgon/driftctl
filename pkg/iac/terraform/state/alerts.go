package state

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type StateReadingAlert struct {
	key string
	err string
}

func NewStateReadingAlert(key string, err error) *StateReadingAlert {
	_ = "STUB: not implemented"
	return nil
}

func (s *StateReadingAlert) Message() string { _ = "STUB: not implemented"; return "" }

func (s *StateReadingAlert) ShouldIgnoreResource() bool { _ = "STUB: not implemented"; return false }

func (s *StateReadingAlert) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

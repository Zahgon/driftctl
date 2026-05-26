package alerter

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type Alerts map[string][]Alert

type Alert interface {
	Message() string
	ShouldIgnoreResource() bool
	Resource() *resource.Resource
}

type UnsupportedResourcetypeAlert struct {
	Typ string
}

func NewUnsupportedResourcetypeAlert(typ string) *UnsupportedResourcetypeAlert {
	_ = "STUB: not implemented"
	return nil
}

func (f *UnsupportedResourcetypeAlert) Message() string { _ = "STUB: not implemented"; return "" }

func (f *UnsupportedResourcetypeAlert) ShouldIgnoreResource() bool {
	_ = "STUB: not implemented"
	return false
}

func (f *UnsupportedResourcetypeAlert) Resource() *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

type FakeAlert struct {
	Msg            string
	IgnoreResource bool
}

func (f *FakeAlert) Message() string { _ = "STUB: not implemented"; return "" }

func (f *FakeAlert) ShouldIgnoreResource() bool { _ = "STUB: not implemented"; return false }

func (f *FakeAlert) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

type SerializableAlert struct {
	Alert
}

type SerializedAlert struct {
	Msg string `json:"message"`
}

func (u *SerializedAlert) Message() string { _ = "STUB: not implemented"; return "" }

func (u *SerializedAlert) ShouldIgnoreResource() bool { _ = "STUB: not implemented"; return false }

func (s *SerializedAlert) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

func (s *SerializableAlert) UnmarshalJSON(bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SerializableAlert) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

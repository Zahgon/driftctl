package common

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type Enumerator interface {
	SupportedType() resource.ResourceType
	Enumerate() ([]*resource.Resource, error)
}

type RemoteLibrary struct {
	enumerators []Enumerator
}

func NewRemoteLibrary() *RemoteLibrary { _ = "STUB: not implemented"; return nil }

func (r *RemoteLibrary) AddEnumerator(enumerator Enumerator) { _ = "STUB: not implemented"; return }

func (r *RemoteLibrary) Enumerators() []Enumerator { _ = "STUB: not implemented"; return nil }

package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type Chain []Middleware

func NewChain(middlewares ...Middleware) Chain { _ = "STUB: not implemented"; return *new(Chain) }

func (c Chain) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

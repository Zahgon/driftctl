package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Manage tags_all attribute on each compatible resources
type TagsAllManager struct{}

func NewTagsAllManager() TagsAllManager { _ = "STUB: not implemented"; return *new(TagsAllManager) }

func (a TagsAllManager) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

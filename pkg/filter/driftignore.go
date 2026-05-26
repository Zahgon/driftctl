package filter

import (
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/snyk/driftctl/enumeration/resource"
)

const separator = "_-_"

type DriftIgnore struct {
	driftignorePath string
	ignorePatterns  []string
	matcher         gitignore.Matcher
}

func NewDriftIgnore(path string, ignorePatterns ...string) *DriftIgnore {
	_ = "STUB: not implemented"
	return nil
}

func (r *DriftIgnore) readIgnoreFile() error { _ = "STUB: not implemented"; return nil }

func (r *DriftIgnore) parseIgnorePatterns() error { _ = "STUB: not implemented"; return nil }

func (r *DriftIgnore) parseIgnorePattern(line string, patterns *[]gitignore.Pattern) {
	_ = "STUB: not implemented"
	return
}

// empty

// this is a comment

func (r *DriftIgnore) isAnyOfChildrenTypesNotIgnored(ty resource.ResourceType) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *DriftIgnore) IsTypeIgnored(ty resource.ResourceType) bool {
	_ = "STUB: not implemented"
	// Iterate over children types, and do not ignore parent resource
	// if at least one of children type is not ignored.
	return false
}

func (r *DriftIgnore) shouldIgnoreType(ty resource.ResourceType) bool {
	_ = "STUB: not implemented"
	return false
}

// If a line start with a `!` and if the type match, we should not ignore it

func (r *DriftIgnore) IsResourceIgnored(res *resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *DriftIgnore) match(strRes string) bool { _ = "STUB: not implemented"; return false }

package test

import (
	"testing"

	"github.com/snyk/driftctl/enumeration/terraform"

	"github.com/snyk/driftctl/enumeration/resource"

	"github.com/r3labs/diff/v2"
)

func doTestDiff(got []*resource.Resource, dirName string, provider terraform.TerraformProvider, deserializer *resource.Deserializer, shouldUpdate bool) (diff.Changelog, error) {
	_ = "STUB: not implemented"
	return *new(diff.Changelog), nil
}

// CtyTestDiff Deprecated
func CtyTestDiff(got []*resource.Resource, dirName string, provider terraform.TerraformProvider, deserializer *resource.Deserializer, shouldUpdate bool, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

package test

import (
	"testing"

	"github.com/snyk/driftctl/enumeration/terraform"

	"github.com/snyk/driftctl/enumeration/resource"
)

// That method is used to compare the result of the enumeration with the golden file.
// That method does not use cty and types from the terraform provider to deserialize resources.
// Some resources returned by the enumeration may have missing fields, and if we use cty deserialization we're
// gonna recreate those missing fields to respect the schema.
func TestAgainstGoldenFileNoCty(
	got []*resource.Resource,
	ty string,
	dirName string,
	_ terraform.TerraformProvider,
	_ *resource.Deserializer,
	shouldUpdate bool,
	tt *testing.T) {
	_ = "STUB: not implemented"
	return
}

// update golden file

// read golden file

// diff

func testAgainstGoldenFileCty(
	got []*resource.Resource,
	ty string,
	dirName string,
	provider terraform.TerraformProvider,
	deserializer *resource.Deserializer,
	shouldUpdate bool,
	tt *testing.T,
) {
	_ = "STUB: not implemented"
	return
}

// update golden file

// read golden file

// diff

func TestAgainstGoldenFile(
	got []*resource.Resource,
	ty string,
	dirName string,
	provider terraform.TerraformProvider,
	deserializer *resource.Deserializer,
	shouldUpdate bool,
	tt *testing.T,
) {
	_ = "STUB: not implemented"
	return
}

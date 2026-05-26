package schemas

import (
	"embed"

	"github.com/hashicorp/terraform/providers"
)

//go:embed */*/schema.json
var fakeSchemaFS embed.FS

func WriteTestSchema(schema map[string]providers.Schema, provider, version string) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadTestSchema(provider, version string) (map[string]providers.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

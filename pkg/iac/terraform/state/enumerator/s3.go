package enumerator

import (
	"github.com/aws/aws-sdk-go/service/s3/s3iface"

	"github.com/snyk/driftctl/pkg/iac/config"
)

type S3Enumerator struct {
	config config.SupplierConfig
	client s3iface.S3API
}

func NewS3Enumerator(config config.SupplierConfig) *S3Enumerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *S3Enumerator) Origin() string { _ = "STUB: not implemented"; return "" }

func (s *S3Enumerator) Enumerate() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// prefix should contains everything that does not have a glob pattern
// Pattern should be the glob matcher string

// We combine the prefix and pattern to match file names against.

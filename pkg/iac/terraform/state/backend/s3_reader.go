package backend

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

const BackendKeyS3 = "s3"

type S3Backend struct {
	input    s3.GetObjectInput
	reader   io.ReadCloser
	S3Client s3iface.S3API
}

func NewS3Reader(path string) (*S3Backend, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *S3Backend) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *S3Backend) Close() error { _ = "STUB: not implemented"; return nil }

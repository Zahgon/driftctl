package backend

import (
	"io"

	"cloud.google.com/go/storage"
)

const BackendKeyGS = "gs"

type GSBackend struct {
	bucketName    string
	path          string
	reader        io.ReadCloser
	storageClient *storage.Client
}

func NewGSReader(path string) (*GSBackend, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *GSBackend) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *GSBackend) Close() error { _ = "STUB: not implemented"; return nil }

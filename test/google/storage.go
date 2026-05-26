package google

import (
	"net/http"
	"net/http/httptest"

	"cloud.google.com/go/storage"
)

type FakeStorageServer struct {
	routes map[string]http.HandlerFunc
}

func (s *FakeStorageServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func NewFakeStorageServer(routes map[string]http.HandlerFunc) (*storage.Client, *httptest.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func newStorageClient(fakeServer *FakeStorageServer) (*storage.Client, *httptest.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

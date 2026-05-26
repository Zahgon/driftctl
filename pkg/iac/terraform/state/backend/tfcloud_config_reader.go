package backend

import (
	"io"
)

type container struct {
	Credentials map[string]containerToken
}

type containerToken struct {
	Token string
}

type tfCloudConfigReader struct {
	reader io.ReadCloser
}

func NewTFCloudConfigReader(reader io.ReadCloser) *tfCloudConfigReader {
	_ = "STUB: not implemented"
	return nil
}

func (r *tfCloudConfigReader) GetToken(host string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getTerraformConfigFile() (string, error) { _ = "STUB: not implemented"; return "", nil }

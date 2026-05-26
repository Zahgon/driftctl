package backend

import (
	"io"
	"regexp"

	tfe "github.com/hashicorp/go-tfe"
)

const BackendKeyTFCloud = "tfcloud"

type TFCloudAttributes struct {
	HostedStateDownloadUrl string `json:"hosted-state-download-url"`
}

type TFCloudData struct {
	Attributes TFCloudAttributes `json:"attributes"`
}

type TFCloudBody struct {
	Data TFCloudData `json:"data"`
}

type TFCloudBackend struct {
	client        *tfe.Client
	reader        io.ReadCloser
	opts          *Options
	workspacePath string
}

func NewTFCloudReader(workspacePath string, opts *Options) *TFCloudBackend {
	_ = "STUB: not implemented"
	return nil
}

func (t *TFCloudBackend) getToken() (string, error) { _ = "STUB: not implemented"; return "", nil }

// A regular expression used to validate string workspace ID patterns.
var reStringID = regexp.MustCompile(`^ws-[a-zA-Z0-9\-\._]+$`)

// isValidWorkspaceID checks if the given input is present and non-empty.
func isValidWorkspaceID(v string) bool { _ = "STUB: not implemented"; return false }

func (t *TFCloudBackend) getWorkspaceId() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *TFCloudBackend) initTFEClient() error { _ = "STUB: not implemented"; return nil }

func (t *TFCloudBackend) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t *TFCloudBackend) Close() error { _ = "STUB: not implemented"; return nil }

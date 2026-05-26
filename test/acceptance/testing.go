package acceptance

import (
	"testing"
	"time"

	"github.com/snyk/driftctl/test"

	"github.com/spf13/cobra"

	"github.com/snyk/driftctl/pkg/cmd"

	"github.com/hashicorp/terraform-exec/tfexec"
)

type ShouldRetryFunc func(result *test.ScanResult, retryDuration time.Duration, retryCount uint8) bool

type AccCheck struct {
	PreExec     func()
	PostExec    func()
	Env         map[string]string
	Args        func() []string
	ShouldRetry ShouldRetryFunc
	Check       func(result *test.ScanResult, stdout string, err error)
}

type AccTestCase struct {
	DoNotRunTerraform          bool
	TerraformVersion           string
	WorkingDir                 string
	Paths                      []string
	Args                       []string
	OnStart                    func()
	OnEnd                      func()
	Checks                     []AccCheck
	tmpResultFilePath          string
	originalEnv                []string
	tf                         map[string]*tfexec.Terraform
	ShouldRefreshBeforeDestroy bool
}

func (c *AccTestCase) initTerraformExecutor() error { _ = "STUB: not implemented"; return nil }

func (c *AccTestCase) createResultFile(t *testing.T) error { _ = "STUB: not implemented"; return nil }

func (c *AccTestCase) validate() error { _ = "STUB: not implemented"; return nil }

func (c *AccTestCase) getResultFilePath() string { _ = "STUB: not implemented"; return "" }

func (c *AccTestCase) getResult(t *testing.T) *test.ScanResult {
	_ = "STUB: not implemented"
	return nil
}

/**
 * Retrieve env from os.Environ() but override every variable prefixed with ACC_
 * e.g. ACC_AWS_PROFILE will override AWS_PROFILE
 */
func (c *AccTestCase) resolveTerraformEnv() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *AccTestCase) terraformInit() error { _ = "STUB: not implemented"; return nil }

func (c *AccTestCase) terraformApply() error { _ = "STUB: not implemented"; return nil }

func (c *AccTestCase) terraformDestroy() error { _ = "STUB: not implemented"; return nil }

func (c *AccTestCase) terraformRefresh() error { _ = "STUB: not implemented"; return nil }

func runDriftCtlCmd(driftctlCmd *cmd.DriftctlCmd) (*cobra.Command, string, error) {
	_ = "STUB: not implemented"
	// keep backup of the real stdout
	return nil, "", nil
}

// Ignore not in sync errors in acceptance test context

// copy the output in a separate goroutine so printing can't block indefinitely

// back to normal state

// restoring the real stdout

func (c *AccTestCase) useTerraformEnv() { _ = "STUB: not implemented"; return }

func (c *AccTestCase) restoreEnv() { _ = "STUB: not implemented"; return }

func (c *AccTestCase) setEnv(env []string) { _ = "STUB: not implemented"; return }

func Run(t *testing.T, c AccTestCase) { _ = "STUB: not implemented"; return }

// Disable terraform version checks
// @link https://www.terraform.io/docs/commands/index.html#upgrade-and-security-bulletin-checks

// Retry after 2s, 4s, 8s, 16s, 32s, 64s, 2m, 2m, 2m, 2m
// Try tweaking the backoff interval limit and/or the retry count limit in
// response to flaky tests.

// Execute terraform init if .terraform folder is not found in test folder

// If the path contains only one element, we switch to this directory for driftctl execution
// We can override this logic by passing a WorkingDir argument in test

// If any --from flag was manually provided OR if a working dir is specified,
// do not setup any --from flags

// Restore original working directory

// LinearBackoff returns a function that retries using
// a back-off strategy of retrying 'n' times and doubling the
// amount of time waited after each one.
func LinearBackoff(limit time.Duration) ShouldRetryFunc {
	_ = "STUB: not implemented"
	return *new(ShouldRetryFunc)
}

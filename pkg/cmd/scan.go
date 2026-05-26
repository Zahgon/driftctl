package cmd

import (
	"github.com/snyk/driftctl/pkg/iac/config"
	"github.com/spf13/cobra"

	"github.com/snyk/driftctl/pkg"
)

func NewScanCmd(opts *pkg.ScanOptions) *cobra.Command { _ = "STUB: not implemented"; return nil }

// Attempt to read the provider version from a terraform lock file

func scanRun(opts *pkg.ScanOptions) error { _ = "STUB: not implemented"; return nil }

// For now, we only use the global printer to print progress and information about the current scan, so unless one
// of the configured output should silence global output we simply use console by default.

// special case command-line advice, because AWS is the default cloud
// provider, and users may be confused by a cloud-specific error out of
// the box

// Teardown

// TODO use enum library interface here

// Fallback to console output if all output failed

func validateTfProviderVersionString(version string) error { _ = "STUB: not implemented"; return nil }

func retrieveBackendsFromHCL(workdir string) ([]config.SupplierConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

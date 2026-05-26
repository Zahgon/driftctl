package cmd

import (
	"github.com/snyk/driftctl/build"
	"github.com/spf13/cobra"
)

var usageTemplate = `Usage: {{.UseLine}}{{if .HasAvailableSubCommands}}

COMMANDS:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name 24 }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

FLAGS:
{{ .LocalFlags.FlagUsages | trimTrailingWhitespaces }}{{end}}{{if .HasAvailableInheritedFlags}}

INHERITED FLAGS:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableSubCommands}}

LEARN MORE:
  Use "{{.CommandPath}} <command> --help" for more information about a command{{end}}
`

var versionTemplate = `{{ printf "%s\n" .Version }}`

type DriftctlCmd struct {
	cobra.Command
	build build.BuildInterface
}

func NewDriftctlCmd(build build.BuildInterface) *DriftctlCmd { _ = "STUB: not implemented"; return nil }

func contains(args []string, cmd string) bool { _ = "STUB: not implemented"; return false }

func (driftctlCmd DriftctlCmd) ShouldCheckVersion() bool { _ = "STUB: not implemented"; return false }

func IsReportingEnabled(cmd *cobra.Command) bool { _ = "STUB: not implemented"; return false }

func handleReporting(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

// Iterate over command flags
// If the command flag is not manually set (f.Changed) we override its value
// from the according env value
func bindEnvToFlags(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

// Ignore some global flags
// no-version-check is ignored because we don't use cmd flags to retrieve flag in version check function
// as we check version before cmd, we use os.Args

// Apply the viper config value to the flag when the flag is not set and viper has a value
// Allow flags precedence over env variables

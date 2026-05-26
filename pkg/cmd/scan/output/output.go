package output

import (
	"github.com/snyk/driftctl/pkg/analyser"
	"github.com/snyk/driftctl/pkg/output"
)

type Output interface {
	Write(analysis *analyser.Analysis) error
}

var supportedOutputTypes = []string{
	ConsoleOutputType,
	JSONOutputType,
	HTMLOutputType,
	PlanOutputType,
}

var supportedOutputExample = map[string]string{
	ConsoleOutputType: ConsoleOutputExample,
	JSONOutputType:    JSONOutputExample,
	HTMLOutputType:    HTMLOutputExample,
	PlanOutputType:    PlanOutputExample,
}

func SupportedOutputsExample() []string { _ = "STUB: not implemented"; return nil }

func Example(key string) string { _ = "STUB: not implemented"; return "" }

func IsSupported(key string) bool { _ = "STUB: not implemented"; return false }

func GetOutput(config OutputConfig) Output { _ = "STUB: not implemented"; return *new(Output) }

// ShouldPrint indicate if we should use the global output or not (e.g. when outputting to stdout).
func ShouldPrint(outputs []OutputConfig, quiet bool) bool { _ = "STUB: not implemented"; return false }

func GetPrinter(config OutputConfig, quiet bool) output.Printer {
	_ = "STUB: not implemented"
	return *new(output.Printer)
}

func isStdOut(path string) bool { _ = "STUB: not implemented"; return false }

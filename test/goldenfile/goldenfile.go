package goldenfile

import (
	"flag"
	"path/filepath"
	"runtime"
)

const (
	GoldenFilePath  = "test"
	ResultsFilename = "results.golden.json"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)

var Update = flag.String("update", "", "name of test to update")

func ReadRootFile(p string, name string) []byte { _ = "STUB: not implemented"; return nil }

func WriteRootFile(p string, content []byte, name string) {
	_ = "STUB: not implemented"

	// Avoid creating golden files for empty results
	return
}

func ReadFile(p string, name string) []byte { _ = "STUB: not implemented"; return nil }

func WriteFile(p string, content []byte, name string) {
	_ = "STUB: not implemented"

	// Avoid creating golden files for empty results
	return
}

// Remove forbidden characters like / in file name
func sanitizeName(name string) string { _ = "STUB: not implemented"; return "" }

func FileExists(dirname, f string) bool { _ = "STUB: not implemented"; return false }

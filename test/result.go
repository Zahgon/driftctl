package test

import (
	"testing"

	"github.com/snyk/driftctl/pkg/analyser"

	"github.com/stretchr/testify/require"
)

type ScanResult struct {
	*require.Assertions
	*analyser.Analysis
}

func NewScanResult(t *testing.T, analysis *analyser.Analysis) *ScanResult {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScanResult) AssertResourceUnmanaged(id, ty string) { _ = "STUB: not implemented"; return }

func (r *ScanResult) AssertResourceDeleted(id, ty string) { _ = "STUB: not implemented"; return }

func (r *ScanResult) AssertCoverage(expected int) { _ = "STUB: not implemented"; return }

func (r *ScanResult) AssertDeletedCount(count int) { _ = "STUB: not implemented"; return }

func (r *ScanResult) AssertManagedCount(count int) { _ = "STUB: not implemented"; return }

func (r *ScanResult) AssertUnmanagedCount(count int) { _ = "STUB: not implemented"; return }

func (r ScanResult) AssertInfrastructureIsInSync() { _ = "STUB: not implemented"; return }

func (r ScanResult) AssertInfrastructureIsNotSync() { _ = "STUB: not implemented"; return }

func (r *ScanResult) printAnalysisResult() string { _ = "STUB: not implemented"; return "" }

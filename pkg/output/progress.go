package output

import (
	"time"

	"go.uber.org/atomic"
)

var spinner = []string{"⣷", "⣯", "⣟", "⡿", "⢿", "⣻", "⣽", "⣾"}

const (
	progressTimeout     = 10 * time.Second
	progressRefreshRate = 200 * time.Millisecond
)

type Progress interface {
	Start()
	Stop()
	Inc()
	Val() uint64
}

type ProgressOptions struct {
	LoadingText  string
	FinishedText string
	ShowCount    bool
}

type progress struct {
	endChan           chan struct{}
	started           *atomic.Bool
	count             *atomic.Uint64
	loadingText       string
	finishedText      string
	showCount         bool
	highestLineLength int
}

func NewProgress(loadingText, finishedText string, showCount bool) *progress {
	_ = "STUB: not implemented"
	return nil
}

func (p *progress) Start() { _ = "STUB: not implemented"; return }

func (p *progress) Stop() { _ = "STUB: not implemented"; return }

func (p *progress) Inc() { _ = "STUB: not implemented"; return }

func (p *progress) Val() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *progress) render() { _ = "STUB: not implemented"; return }

func (p *progress) watch() { _ = "STUB: not implemented"; return }

func (p *progress) flush() { _ = "STUB: not implemented"; return }

func (p *progress) printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

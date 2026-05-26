package parallel

import (
	"context"
	"sync"

	"go.uber.org/atomic"

	"golang.org/x/sync/semaphore"
)

type ParallelRunner struct {
	sem     *semaphore.Weighted
	wg      *sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
	resChan chan interface{}
	err     error
	hasErr  *atomic.Bool
	waiting *atomic.Bool
}

func NewParallelRunner(ctx context.Context, maxRun int64) *ParallelRunner {
	_ = "STUB: not implemented"
	return nil
}

func (p *ParallelRunner) SubRunner() *ParallelRunner { _ = "STUB: not implemented"; return nil }

func (p *ParallelRunner) Read() chan interface{} { _ = "STUB: not implemented"; return nil }

func (p *ParallelRunner) DoneChan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (p *ParallelRunner) Err() error { _ = "STUB: not implemented"; return nil }

func (p *ParallelRunner) wait() { _ = "STUB: not implemented"; return }

func (p *ParallelRunner) Run(runnable func() (interface{}, error)) {
	_ = "STUB: not implemented"
	return
}

// only release if sem was acquired

// Prevent new routines executions if we already got an error from another routine

// Handle panic in routines and stop runner with proper error
// Some failed call to grpc plugin like getSchema trigger a panic

func (p *ParallelRunner) Stop(err error) { _ = "STUB: not implemented"; return }

package sentry

import (
	cmderrors "github.com/snyk/driftctl/pkg/cmd/errors"
)

var excludedErrorTypes = []error{
	cmderrors.UsageError{},
}

func Initialize() error { _ = "STUB: not implemented"; return nil }

func shouldCaptureException(err error) bool { _ = "STUB: not implemented"; return false }

func CaptureException(err error) { _ = "STUB: not implemented"; return }

package errors

type UsageError struct {
	msg string
}

func NewUsageError(msg string) UsageError { _ = "STUB: not implemented"; return *new(UsageError) }

func (u UsageError) Error() string { _ = "STUB: not implemented"; return "" }

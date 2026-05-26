package iac

type StateReadingError struct {
	errors []error
}

func NewStateReadingError() *StateReadingError { _ = "STUB: not implemented"; return nil }

func (s *StateReadingError) Add(err error) { _ = "STUB: not implemented"; return }

func (s *StateReadingError) Error() string { _ = "STUB: not implemented"; return "" }

package errors

type InfrastructureNotInSync struct{}

func (i InfrastructureNotInSync) Error() string { _ = "STUB: not implemented"; return "" }

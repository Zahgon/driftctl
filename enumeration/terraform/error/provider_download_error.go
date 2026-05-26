package error

type ProviderNotFoundError struct {
	Version string
}

func (p ProviderNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

package error

type RemoteError interface {
	ListedTypeError() string
}

type ResourceScanningError struct {
	err             error
	resourceType    string
	resourceId      string
	listedTypeError string
}

func (b *ResourceScanningError) Error() string { _ = "STUB: not implemented"; return "" }

func (b *ResourceScanningError) RootCause() error { _ = "STUB: not implemented"; return nil }

func (b *ResourceScanningError) ResourceType() string { _ = "STUB: not implemented"; return "" }

func NewResourceScanningError(error error, resourceType string, resourceId string) *ResourceScanningError {
	_ = "STUB: not implemented"
	return nil
}

func NewResourceListingError(error error, resourceType string) *ResourceScanningError {
	_ = "STUB: not implemented"
	return nil
}

func NewResourceListingErrorWithType(error error, resourceType, listedTypeError string) *ResourceScanningError {
	_ = "STUB: not implemented"
	return nil
}

func (b *ResourceScanningError) ListedTypeError() string { _ = "STUB: not implemented"; return "" }

func (b *ResourceScanningError) Resource() string { _ = "STUB: not implemented"; return "" }

func (b *ResourceScanningError) String() string { _ = "STUB: not implemented"; return "" }

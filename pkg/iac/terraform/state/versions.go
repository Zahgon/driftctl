package state

import (
	"github.com/hashicorp/go-version"
)

var (
	// UnsupportedVersionConstraints is an array of version constraints known to be unsupported.
	// If a given state matches one of these, all resources of the related state will be ignored and marked as drifted.
	UnsupportedVersionConstraints = []string{"<0.11.0"}
)

type UnsupportedVersionError struct {
	StateFile string
	Version   *version.Version
}

func (u *UnsupportedVersionError) Error() string { _ = "STUB: not implemented"; return "" }

func IsVersionSupported(rawVersion string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

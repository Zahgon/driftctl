package version

// Current software version
// Could be injected on build with -ldflags
var version string = "dev"

// Return the current version as string
func Current() string { _ = "STUB: not implemented"; return "" }

/**
 * Return "" if current version is the last version,
 * else return the latest version string
 **/
func CheckLatest() string { _ = "STUB: not implemented"; return "" }

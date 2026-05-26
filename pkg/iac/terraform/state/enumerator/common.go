package enumerator

// Returns the below segments:
// - prefix : path part that should not contains glob patterns, that is used in S3 query to filter result
// - pattern : should contains the glob pattern to be used by doublestar matching library
func extractPrefixAndPattern(path string) (prefix string, pattern string) {
	_ = "STUB: not implemented"
	return "", ""
}

// HasMeta reports whether path contains any of the magic characters
func HasMeta(path string) bool { _ = "STUB: not implemented"; return false }

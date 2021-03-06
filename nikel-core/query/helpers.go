package query

import (
	"strconv"
	"strings"
)

// ParseInt is a light wrapper around strconv.ParseInt with bound checks
func ParseInt(query string, low, high, def int) int {
	// don't bother to do anything if query is empty
	if query == "" {
		return def
	}

	// parse int on base 10 with size 64bit
	i, err := strconv.ParseInt(query, 10, 64)

	// do necessary checks to make sure value is valid
	if val := int(i); err == nil && low <= val && val <= high {
		return val
	}

	return def
}

// containsCaseInsensitive is the same as contains but is case insensitive
// it's slow but it is slightly faster than the regex implementation
// this could be further optimized
func containsCaseInsensitive(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

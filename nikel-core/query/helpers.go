package query

import (
	"strconv"
)

// ParseInt is a light wrapper around strconv.ParseInt with bound checks
func ParseInt(query string, low, high, def int) int {
	if query == "" {
		return def
	}
	i, err := strconv.ParseInt(query, 10, 64)
	if val := int(i); err != nil || val < low || val > high {
		return def
	} else {
		return val
	}
}

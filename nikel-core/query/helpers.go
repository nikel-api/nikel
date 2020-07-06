package query

import (
	"strconv"
)

// ParseInt is a light wrapper around ParseInt with bound checks
func ParseInt(query string, low, high, def int) int {
	if query == "" {
		return def
	}
	i, err := strconv.ParseInt(query, 10, 64)
	val := int(i)
	if err != nil || val < low || val > high {
		return def
	}
	return val
}

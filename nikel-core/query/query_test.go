package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPrefixHandlerTableDriven tests all handlers
func TestPrefixHandlerTableDriven(t *testing.T) {
	// defined tests table
	var tests = []struct {
		query        string
		prefix, rest string
	}{
		{"<=test", "<=", "test"},
		{">=test", ">=", "test"},
		{"<test", "<", "test"},
		{">test", ">", "test"},
		{"(test", "startsWith", "test"},
		{")test", "endsWith", "test"},
		{"=test", "=", "test"},
		{"!test", "!=", "test"},
		{"~test", "interface", "test"},
		{"test", "default", "test"},
		{"<=", "<=", ""},
		{"<", "<", ""},
		{"", "default", ""},
	}

	// run test table
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			prefix, rest := prefixHandler(tt.query)
			assert.Equal(t, tt.prefix, prefix)
			assert.Equal(t, tt.rest, rest)
		})
	}
}

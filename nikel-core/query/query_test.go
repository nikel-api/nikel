package query

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPrefixHandlerTableDriven tests all handlers
func TestPrefixHandlerTableDriven(t *testing.T) {
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
		{"test", "default", "test"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.query)
		t.Run(testname, func(t *testing.T) {
			prefix, rest := PrefixHandler(tt.query)
			assert.Equal(t, tt.prefix, prefix)
			assert.Equal(t, tt.rest, rest)
		})
	}
}

// TestTypeToOpTableDriven tests a range of type to op scenarios
func TestTypeToOpTableDriven(t *testing.T) {
	var tests = []struct {
		valueType, op string
		want          string
	}{
		{"string", "!=", "!="},
		{"string", "default", "contains"},
		{"nonString", "!=", "!="},
		{"nonString", "default", "="},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("(%s, %s)", tt.valueType, tt.op)
		t.Run(testname, func(t *testing.T) {
			want := TypeToOp(tt.valueType, tt.op)
			assert.Equal(t, tt.want, want)
		})
	}
}

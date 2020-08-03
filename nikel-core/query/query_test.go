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
		{"~test", "interface", "test"},
		{"test", "default", "test"},
		{"<=", "<=", ""},
		{"<", "<", ""},
		{"", "default", ""},
	}
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			prefix, rest := prefixHandler(tt.query)
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
			want := typeToOp(tt.valueType, tt.op)
			assert.Equal(t, tt.want, want)
		})
	}
}

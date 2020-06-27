// TODO: AutoQuery Tests
package query

import (
	"fmt"
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
			if prefix != tt.prefix || rest != tt.rest {
				t.Errorf("got (%s, %s), want (%s, %s)", prefix, rest, tt.prefix, tt.rest)
			}
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
			if want != tt.want {
				t.Errorf("got %s, want %s", want, tt.want)
			}
		})
	}
}

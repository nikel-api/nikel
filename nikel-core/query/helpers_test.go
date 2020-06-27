package query

import "testing"

// TestParseIntQuery tests if query gets parsed
func TestParseIntQuery(t *testing.T) {
	res := ParseInt("50", 0, 100, 10)
	if res != 50 {
		t.Errorf(`ParseInt("50", 0, 100, 10) = %d; want 50`, res)
	}
}

// TestParseIntLow tests if low is triggered
func TestParseIntLow(t *testing.T) {
	res := ParseInt("-10", 0, 100, 10)
	if res != 10 {
		t.Errorf(`ParseInt("-10", 0, 100, 10) = %d; want 10`, res)
	}
}

// TestParseIntHigh tests if high is triggered
func TestParseIntHigh(t *testing.T) {
	res := ParseInt("110", 0, 100, 10)
	if res != 10 {
		t.Errorf(`ParseInt("110", 0, 100, 10) = %d; want 10`, res)
	}
}

package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestConfig basically tests nothing important
func TestConfig(t *testing.T) {
	assert.Greater(t, DefaultLimit, 0)
	assert.IsType(t, DefaultLimit, 0)
	assert.Greater(t, TopLimit, DefaultLimit)
	assert.IsType(t, TopLimit, DefaultLimit)
}

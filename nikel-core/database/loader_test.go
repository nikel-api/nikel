package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/thedevsaddam/gojsonq/v2"
	"reflect"
	"testing"
)

// TestLoadDatabase tests database loader
func TestLoadDatabase(t *testing.T) {
	// this is some super ugly reflect code, but it's fine since it's only used in test code
	db := reflect.ValueOf(*DB)
	for i := 0; i < db.NumField(); i++ {
		assert.NotZero(t, db.Field(i).Interface().(*gojsonq.JSONQ).Count())
	}
}

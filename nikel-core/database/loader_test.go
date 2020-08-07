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

	// loop through all the struct fields
	for i := 0; i < db.NumField(); i++ {
		// get struct field and count the elements loaded
		// should be non zero (no way it's negative right?)
		assert.NotZero(t, db.Field(i).Interface().(*gojsonq.JSONQ).Count())
	}
}

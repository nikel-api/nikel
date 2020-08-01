package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestLoadDatabase tests database loader
func TestLoadDatabase(t *testing.T) {
	for _, val := range []int{
		DB.BuildingsData.Count(),
		DB.CoursesData.Count(),
		DB.EvalsData.Count(),
		DB.ExamsData.Count(),
		DB.FoodData.Count(),
		DB.ParkingData.Count(),
		DB.ProgramsData.Count(),
		DB.ServicesData.Count(),
		DB.TextbooksData.Count(),
	} {
		assert.Greater(t, val, 0)
	}
}

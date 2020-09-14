package database

import (
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/query"
	"github.com/thedevsaddam/gojsonq/v2"
)

// LoadFile loads file
func LoadFile(path string) *gojsonq.JSONQ {
	// use Reset to force a GC run on raw string data inside struct
	jq := gojsonq.New().File(config.PathPrefix + path).Reset()
	jq.Macro("interface", query.InterfaceMacro)
	return jq
}

// init loads JSON data to database
func init() {
	// load database
	// seriously considering a non hardcoded reflection solution
	// because this really doesn't look right
	DB.CoursesData = LoadFile(config.CoursesPath)
	DB.ProgramsData = LoadFile(config.ProgramsPath)
	DB.TextbooksData = LoadFile(config.TextbooksPath)
	DB.BuildingsData = LoadFile(config.BuildingsPath)
	DB.FoodData = LoadFile(config.FoodPath)
	DB.ParkingData = LoadFile(config.ParkingPath)
	DB.ServicesData = LoadFile(config.ServicesPath)
	DB.ExamsData = LoadFile(config.ExamsPath)
	DB.EvalsData = LoadFile(config.EvalsPath)
}

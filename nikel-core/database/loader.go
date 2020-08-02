package database

import (
	"fmt"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/query"
	"github.com/thedevsaddam/gojsonq/v2"
	"os"
	"path/filepath"
)

var pathPrefix = ""

// LoadFile loads file
func LoadFile(path string) *gojsonq.JSONQ {
	// Use Reset to force a GC run on raw string data inside struct
	jq := gojsonq.New().File(pathPrefix + path).Reset()
	jq.Macro("interface", query.InterfaceMacro)
	return jq
}

// init loads JSON data to database
func init() {
	wd, _ := os.Getwd()

	// travel up the parent folders to find proper directory position
	steps := 0

	for {
		// exit if travelled up too far
		if steps == 5 {
			panic(fmt.Errorf("nikel-core: cannot find folder positions"))
		}

		// find go.mod file
		if _, err := os.Stat(fmt.Sprintf("%s/%s", wd, "go.mod")); !os.IsNotExist(err) {
			break
		}

		pathPrefix += "../"
		wd = filepath.Dir(wd)
		steps += 1
	}

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

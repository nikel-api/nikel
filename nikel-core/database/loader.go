package database

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/query"
	"github.com/thedevsaddam/gojsonq/v2"
	"os"
	"path/filepath"
)

var pathPrefix = ""

type decoder struct{}

// Decode is a decode wrapper around jsoniter
func (d *decoder) Decode(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, &v)
}

// loadFile loads file
func loadFile(path string) *gojsonq.JSONQ {
	// Use Reset to force a GC run on raw string data inside struct
	jq := gojsonq.New(gojsonq.SetDecoder(&decoder{})).File(pathPrefix + path).Reset()
	jq.Macro("interface", query.InterfaceMacro)
	return jq
}

// init loads JSON data to database
func init() {
	wd, _ := os.Getwd()

	// travel up the parent folders to find proper directory position
	steps := 0

	// app folder name is for heroku deployment
	for filepath.Base(wd) != "nikel" && filepath.Base(wd) != "app" {

		// exit if travelled up too far
		if steps == 5 {
			panic(fmt.Errorf("nikel-core: cannot find folder positions"))
		}

		pathPrefix += "../"
		wd = filepath.Dir(wd)
		steps += 1
	}

	DB.CoursesData = loadFile(config.CoursePath)
	DB.TextbooksData = loadFile(config.TextbookPath)
	DB.BuildingsData = loadFile(config.BuildingsPath)
	DB.FoodData = loadFile(config.FoodPath)
	DB.ParkingData = loadFile(config.ParkingPath)
	DB.ServicesData = loadFile(config.ServicesPath)
	DB.ExamsData = loadFile(config.ExamsPath)
	DB.EvalsData = loadFile(config.EvalsPath)
}

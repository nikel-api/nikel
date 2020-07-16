package database

import (
	"fmt"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/thedevsaddam/gojsonq/v2"
	"os"
	"path/filepath"
)

// init loads JSON data to database
func init() {
	pathPrefix := ""
	wd, _ := os.Getwd()

	// travel up the parent folders to find proper directory position
	steps := 0

	// app folder name is for heroku deployment
	for filepath.Base(wd) != "nikel" || filepath.Base(wd) != "app" {

		// exit if travelled up too far
		if steps == 5 {
			panic(fmt.Errorf("nikel-core: cannot find folder positions"))
		}

		pathPrefix += "../"
		wd = filepath.Dir(wd)
		steps += 1
	}

	DB.CoursesData = gojsonq.New().File(pathPrefix + config.COURSEPATH).Reset()
	DB.TextbooksData = gojsonq.New().File(pathPrefix + config.TEXTBOOKPATH).Reset()
	DB.BuildingsData = gojsonq.New().File(pathPrefix + config.BUILDINGSPATH).Reset()
	DB.FoodData = gojsonq.New().File(pathPrefix + config.FOODPATH).Reset()
	DB.ParkingData = gojsonq.New().File(pathPrefix + config.PARKINGPATH).Reset()
	DB.ServicesData = gojsonq.New().File(pathPrefix + config.SERVICESPATH).Reset()
	DB.ExamsData = gojsonq.New().File(pathPrefix + config.EXAMSPATH).Reset()
	DB.EvalsData = gojsonq.New().File(pathPrefix + config.EVALSPATH).Reset()
}

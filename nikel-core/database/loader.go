package database

import (
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/thedevsaddam/gojsonq/v2"
	"os"
	"path/filepath"
)

// init loads JSON data to database
func init() {
	pathPrefix := ""
	wd, _ := os.Getwd()

	if filepath.Base(wd) == "nikel-core" {
		pathPrefix = "../"
	} else if filepath.Base(wd) == "router" {
		pathPrefix = "../../"
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

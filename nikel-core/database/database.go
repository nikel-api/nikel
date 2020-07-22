package database

import "github.com/thedevsaddam/gojsonq/v2"

// Database stores Nikel's data
type Database struct {
	CoursesData   *gojsonq.JSONQ
	ProgramsData  *gojsonq.JSONQ
	TextbooksData *gojsonq.JSONQ
	BuildingsData *gojsonq.JSONQ
	FoodData      *gojsonq.JSONQ
	ParkingData   *gojsonq.JSONQ
	ServicesData  *gojsonq.JSONQ
	ExamsData     *gojsonq.JSONQ
	EvalsData     *gojsonq.JSONQ
}

// DB value holds the data
var DB = &Database{}

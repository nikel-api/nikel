package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// contains a variety of constants
// the paths should be removed in the future
const (
	CoursesPath   = "nikel-datasets/data/courses.json"
	ProgramsPath  = "nikel-datasets/data/programs.json"
	TextbooksPath = "nikel-datasets/data/textbooks.json"
	BuildingsPath = "nikel-datasets/data/buildings.json"
	FoodPath      = "nikel-datasets/data/food.json"
	ParkingPath   = "nikel-datasets/data/parking.json"
	ServicesPath  = "nikel-datasets/data/services.json"
	ExamsPath     = "nikel-datasets/data/exams.json"
	EvalsPath     = "nikel-datasets/data/evals.json"
	FaviconPath   = "nikel-core/media/favicon.ico"
	TopLimit      = 100
	DefaultLimit  = 10
)

var PathPrefix = ""

// init calculates PathPrefix
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

		// move up a folder and increment steps
		PathPrefix += "../"
		wd = filepath.Dir(wd)
		steps++
	}
}

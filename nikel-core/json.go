package main

import (
	"encoding/json"
	"gopkg.in/guregu/null.v4"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type Course struct {
	ID                         null.String `json:"id"`
	Code                       null.String `json:"code"`
	Name                       null.String `json:"name"`
	Description                null.String `json:"description"`
	Division                   null.String `json:"division"`
	Department                 null.String `json:"department"`
	Prerequisites              null.String `json:"prerequisites"`
	Corequisites               null.String `json:"corequisites"`
	Exclusions                 null.String `json:"exclusions"`
	RecommendedPreparation     null.String `json:"recommended_preparation"`
	Level                      null.String `json:"level"`
	Campus                     null.String `json:"campus"`
	Term                       null.String `json:"term"`
	ArtsAndScienceBreadth      null.String `json:"arts_and_science_breadth"`
	ArtsAndScienceDistribution null.String `json:"arts_and_science_distribution"`
	UtmDistribution            null.String `json:"utm_distribution"`
	UtscBreadth                null.String `json:"utsc_breadth"`
	ApscElectives              null.String `json:"apsc_electives"`
	MeetingSections            []struct {
		Code        null.String   `json:"code"`
		Instructors []null.String `json:"instructors"`
		Times       []struct {
			Day      null.String `json:"day"`
			Start    null.Int    `json:"start"`
			End      null.Int    `json:"end"`
			Duration null.Int    `json:"duration"`
			Location null.String `json:"location"`
		} `json:"times"`
		Size           null.Int    `json:"size"`
		Enrollment     null.Int    `json:"enrollment"`
		WaitlistOption null.Bool   `json:"waitlist_option"`
		Delivery       null.String `json:"delivery"`
	} `json:"meeting_sections"`
	LastUpdated null.String `json:"last_updated"`
}

type Building struct {
	ID        null.String `json:"id"`
	Code      null.String `json:"code"`
	Tags      null.String `json:"tags"`
	Name      null.String `json:"name"`
	ShortName null.String `json:"short_name"`
	Address   struct {
		Street   null.String `json:"street"`
		City     null.String `json:"city"`
		Province null.String `json:"province"`
		Country  null.String `json:"country"`
		Postal   null.String `json:"postal"`
	} `json:"address"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	LastUpdated null.String `json:"last_updated"`
}

var coursesMap map[string]Course
var coursesOrder []string

var buildingsMap map[string]Building
var buildingsOrder []string

func getByteValue(path string) []byte {
	jsonFile, _ := os.Open(path)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = jsonFile.Close()
	return byteValue
}

func loadVals() {
	pathPrefix := ""
	wd, _ := os.Getwd()
	if filepath.Base(wd) == "nikel-core" {
		pathPrefix = "../"
	}

	_ = json.Unmarshal(getByteValue(pathPrefix+COURSEPATH), &coursesMap)
	for k := range coursesMap {
		coursesOrder = append(coursesOrder, k)
	}
	sort.Strings(coursesOrder)

	_ = json.Unmarshal(getByteValue(pathPrefix+BUILDINGSPATH), &buildingsMap)
	for k := range buildingsMap {
		buildingsOrder = append(buildingsOrder, k)
	}
	sort.Strings(buildingsOrder)
}

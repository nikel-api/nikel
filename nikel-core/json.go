package main

import (
	"github.com/thedevsaddam/gojsonq/v2"
	"gopkg.in/guregu/null.v4"
	"os"
	"path/filepath"
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

type Textbook struct {
	ID      null.String `json:"id"`
	Isbn    null.String `json:"isbn"`
	Title   null.String `json:"title"`
	Edition null.Int    `json:"edition"`
	Author  null.String `json:"author"`
	Image   null.String `json:"image"`
	Price   null.Float  `json:"price"`
	URL     null.String `json:"url"`
	Courses []struct {
		ID              null.String `json:"id"`
		Code            null.String `json:"code"`
		Requirement     null.String `json:"requirement"`
		MeetingSections []struct {
			Code        null.String   `json:"code"`
			Instructors []null.String `json:"instructors"`
		} `json:"meeting_sections"`
	} `json:"courses"`
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

type Food struct {
	ID          null.String `json:"id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	Tags        null.String `json:"tags"`
	Campus      null.String `json:"campus"`
	Address     null.String `json:"address"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	Hours struct {
		Sunday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"sunday"`
		Monday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"monday"`
		Tuesday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"tuesday"`
		Wednesday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"wednesday"`
		Thursday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"thursday"`
		Friday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"friday"`
		Saturday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"saturday"`
	} `json:"hours"`
	Image       null.String   `json:"image"`
	URL         null.String   `json:"url"`
	Twitter     null.String   `json:"twitter"`
	Facebook    null.String   `json:"facebook"`
	Attributes  []null.String `json:"attributes"`
	LastUpdated null.String   `json:"last_updated"`
}

type Parking struct {
	ID          null.String `json:"id"`
	Name        null.String `json:"name"`
	Alias       null.String `json:"alias"`
	BuildingID  null.String `json:"building_id"`
	Description null.String `json:"description"`
	Campus      null.String `json:"campus"`
	Address     null.String `json:"address"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	LastUpdated null.String `json:"last_updated"`
}

type Accessibility struct {
	ID          null.String `json:"id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	BuildingID  null.String `json:"building_id"`
	Campus      null.String `json:"campus"`
	Image       null.String `json:"image"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	Attributes  []null.String `json:"attributes"`
	LastUpdated null.String   `json:"last_updated"`
}

type Database struct {
	CoursesData       *gojsonq.JSONQ
	TextbooksData     *gojsonq.JSONQ
	BuildingsData     *gojsonq.JSONQ
	FoodData          *gojsonq.JSONQ
	ParkingData       *gojsonq.JSONQ
	AccessibilityData *gojsonq.JSONQ
}

var database = &Database{}

func loadVals() {
	pathPrefix := ""
	wd, _ := os.Getwd()
	if filepath.Base(wd) == "nikel-core" {
		pathPrefix = "../"
	}

	database.CoursesData = gojsonq.New().File(pathPrefix + COURSEPATH)
	database.TextbooksData = gojsonq.New().File(pathPrefix + TEXTBOOKPATH)
	database.BuildingsData = gojsonq.New().File(pathPrefix + BUILDINGSPATH)
	database.FoodData = gojsonq.New().File(pathPrefix + FOODPATH)
	database.ParkingData = gojsonq.New().File(pathPrefix + PARKINGPATH)
	database.AccessibilityData = gojsonq.New().File(pathPrefix + ACCESSIBILITYPATH)
}

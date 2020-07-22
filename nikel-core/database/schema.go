package database

import "gopkg.in/guregu/null.v4"

// Course represents a course item
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

// Textbook represents a textbook item
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

// Building represents a building item
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

// Food represents a food item
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

// Parking represents a parking item
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

// Service represents an service item
type Service struct {
	ID          null.String `json:"id"`
	Name        null.String `json:"name"`
	Alias       null.String `json:"alias"`
	BuildingID  null.String `json:"building_id"`
	Description null.String `json:"description"`
	Campus      null.String `json:"campus"`
	Address     null.String `json:"address"`
	Image       null.String `json:"image"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	Tags        null.String   `json:"tags"`
	Attributes  []null.String `json:"attributes"`
	LastUpdated null.String   `json:"last_updated"`
}

// Exam represents an exam item
type Exam struct {
	ID         null.String `json:"id"`
	CourseID   null.String `json:"course_id"`
	CourseCode null.String `json:"course_code"`
	Campus     null.String `json:"campus"`
	Date       null.String `json:"date"`
	Start      null.Int    `json:"start"`
	End        null.Int    `json:"end"`
	Duration   null.Int    `json:"duration"`
	Sections   []struct {
		LectureCode null.String `json:"lecture_code"`
		Split       null.String `json:"split"`
		Location    null.String `json:"location"`
	} `json:"sections"`
	LastUpdated null.String `json:"last_updated"`
}

// Eval represents an eval item
type Eval struct {
	ID     null.String `json:"id"`
	Name   null.String `json:"name"`
	Campus null.String `json:"campus"`
	Terms  []struct {
		Term     null.String `json:"term"`
		Lectures []struct {
			LectureCode null.String `json:"lecture_code"`
			Firstname   null.String `json:"firstname"`
			Lastname    null.String `json:"lastname"`
			S1          null.Float  `json:"s1"`
			S2          null.Float  `json:"s2"`
			S3          null.Float  `json:"s3"`
			S4          null.Float  `json:"s4"`
			S5          null.Float  `json:"s5"`
			S6          null.Float  `json:"s6"`
			Invited     null.Int    `json:"invited"`
			Responses   null.Int    `json:"responses"`
		} `json:"lectures"`
	} `json:"terms"`
	LastUpdated null.String `json:"last_updated"`
}

// Program represents a program item
type Program struct {
	ID          null.String `json:"id"`
	Name        null.String `json:"name"`
	Type        null.String `json:"type"`
	Campus      null.String `json:"campus"`
	Description null.String `json:"description"`
	Enrollment  null.String `json:"enrollment"`
	Completion  null.String `json:"completion"`
	LastUpdated null.String `json:"last_updated"`
}

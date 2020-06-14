package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func getCourses(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(coursesOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Course{}},
		)
		return
	}
	upperLimit := skipQuery + limitQuery
	var courses []Course
	for idx := skipQuery; idx < upperLimit && idx < len(coursesOrder); idx++ {
		courses = append(courses, coursesMap[coursesOrder[idx]])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       courses},
	)
}

func courseHandler(c *gin.Context) {
	p1 := c.Param("p1")
	if p1 == "search" {
		getCoursesBySearch(c)
	} else {
		getCourseByID(c)
	}
}

func getCourseByID(c *gin.Context) {
	id := c.Param("p1")
	if val, ok := coursesMap[id]; ok {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       val},
		)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusNotFound,
			"status_message": "course not found",
			"response":       nil},
		)
	}
}

func getCoursesBySearch(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(coursesOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Course{}},
		)
		return
	}

	// TODO: Optimize performance later
	var resCourses []Course
	for _, v := range coursesOrder {
		course := coursesMap[v]
		if filterQuery(c.Query("id"), course.ID) &&
			filterQuery(c.Query("code"), course.Code) &&
			filterQuery(c.Query("name"), course.Name) &&
			filterQuery(c.Query("description"), course.Description) &&
			filterQuery(c.Query("division"), course.Division) &&
			filterQuery(c.Query("department"), course.Department) &&
			filterQuery(c.Query("prerequisites"), course.Prerequisites) &&
			filterQuery(c.Query("corequisites"), course.Corequisites) &&
			filterQuery(c.Query("exclusions"), course.Exclusions) &&
			filterQuery(c.Query("recommended_preparation"), course.RecommendedPreparation) &&
			filterQuery(c.Query("level"), course.Level) &&
			filterQuery(c.Query("campus"), course.Campus) &&
			filterQuery(c.Query("term"), course.Term) &&
			filterQuery(c.Query("arts_and_science_breadth"), course.ArtsAndScienceBreadth) &&
			filterQuery(c.Query("arts_and_science_distribution"), course.ArtsAndScienceDistribution) &&
			filterQuery(c.Query("utm_distribution"), course.UtmDistribution) &&
			filterQuery(c.Query("utsc_breadth"), course.UtscBreadth) &&
			filterQuery(c.Query("apsc_electives"), course.ApscElectives) {

			for _, w := range course.MeetingSections {
				if filterQuery(c.Query("meeting_code"), w.Code) &&
					filterIntQuery(c.Query("size"), w.Size, 0, math.MaxInt64) &&
					filterIntQuery(c.Query("enrollment"), w.Enrollment, 0, math.MaxInt64) &&
					filterBoolQuery(c.Query("waitlist_option"), w.WaitlistOption) &&
					filterQuery(c.Query("delivery"), w.Delivery) {

					pass := false
					for _, x := range w.Instructors {
						if filterQuery(c.Query("instructor"), x) {
							pass = true
							break
						}
					}

					if pass {
						for _, x := range w.Times {
							if filterQuery(c.Query("day"), x.Day) &&
								filterIntQuery(c.Query("start"), x.Start, 0, 86400) &&
								filterIntQuery(c.Query("end"), x.End, 0, 86400) &&
								filterIntQuery(c.Query("duration"), x.End, 0, 86400) &&
								filterQuery(c.Query("location"), x.Location) {
								resCourses = append(resCourses, course)
								break
							}
						}
					}
				}
			}
		}
	}

	var courses []Course
	upperLimit := skipQuery + limitQuery
	for idx := skipQuery; idx < upperLimit && idx < len(resCourses); idx++ {
		courses = append(courses, resCourses[idx])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       courses},
	)
}

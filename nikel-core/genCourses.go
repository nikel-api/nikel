package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

func getCourses(c *gin.Context) {
	offset := parseInt(c.Query("offset"), 0, math.MaxInt64, 0)
	limit := parseInt(c.Query("limit"), 0, TOPLIMIT, DEFAULTLIMIT)

	data := autoQuery(
		database.CoursesData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var courses []Course
	json.Unmarshal(res, &courses)
	if len(courses) == 0 {
		sendEmptySuccess(c)
	} else {
		sendSuccess(c, courses)
	}
}

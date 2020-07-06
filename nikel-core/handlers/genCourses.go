package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/database"
	"github.com/nikel-api/nikel/nikel-core/query"
	"github.com/nikel-api/nikel/nikel-core/response"
	"math"
)

// GetCourses queries via the courses endpoint
func GetCourses(c *gin.Context) {
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TOPLIMIT, config.DEFAULTLIMIT)

	data := query.AutoQuery(
		database.DB.CoursesData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var courses []database.Course
	json.Unmarshal(res, &courses)
	if len(courses) == 0 {
		response.SendEmptySuccess(c)
	} else {
		response.SendSuccess(c, courses)
	}
}

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

// GetExams queries via the exams endpoint
func GetExams(c *gin.Context) {
	// get offset and limit
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TopLimit, config.DefaultLimit)

	// query data
	data := query.AutoQuery(
		database.DB.ExamsData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	// convert data to bytes
	// this is expensive but reforming the data through an additional unmarshal is necessary
	res, _ := json.Marshal(data)

	// initialize payload variable
	var exams []database.Exam

	// unmarshal on a specific struct defined for this payload
	err := json.Unmarshal(res, &exams)
	if err != nil {
		response.SendError(c)
		return
	}

	if len(exams) == 0 {
		// send empty array
		response.SendEmptySuccess(c)
	} else {
		// send payload
		response.SendSuccess(c, exams)
	}
}

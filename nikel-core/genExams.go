package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

func getExams(c *gin.Context) {
	offset := parseInt(c.Query("offset"), 0, math.MaxInt64, 0)
	limit := parseInt(c.Query("limit"), 0, TOPLIMIT, DEFAULTLIMIT)

	data := autoQuery(
		database.ExamsData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var exams []Exam
	json.Unmarshal(res, &exams)
	if len(exams) == 0 {
		sendEmptySuccess(c)
	} else {
		sendSuccess(c, exams)
	}
}

package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

func getTextbooks(c *gin.Context) {
	offset := parseInt(c.Query("offset"), 0, math.MaxInt64, 0)
	limit := parseInt(c.Query("limit"), 0, TOPLIMIT, DEFAULTLIMIT)

	data := autoQuery(
		database.TextbooksData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var textbooks []Textbook
	json.Unmarshal(res, &textbooks)
	if len(textbooks) == 0 {
		sendEmptySuccess(c)
	} else {
		sendSuccess(c, textbooks)
	}
}

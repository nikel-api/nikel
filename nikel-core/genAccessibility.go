package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

func getAccessibility(c *gin.Context) {
	offset := parseInt(c.Query("offset"), 0, math.MaxInt64, 0)
	limit := parseInt(c.Query("limit"), 0, TOPLIMIT, DEFAULTLIMIT)

	data := autoQuery(
		database.AccessibilityData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var accessibility []Accessibility
	json.Unmarshal(res, &accessibility)
	if len(accessibility) == 0 {
		sendEmptySuccess(c)
	} else {
		sendSuccess(c, accessibility)
	}
}

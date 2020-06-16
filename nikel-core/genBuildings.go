package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

func getBuildings(c *gin.Context) {
	offset := parseInt(c.Query("offset"), 0, math.MaxInt64, 0)
	limit := parseInt(c.Query("limit"), 0, TOPLIMIT, DEFAULTLIMIT)

	data := autoQuery(
		database.BuildingsData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var buildings []Building
	json.Unmarshal(res, &buildings)
	if len(buildings) == 0 {
		sendEmptySuccess(c)
	} else {
		sendSuccess(c, buildings)
	}
}

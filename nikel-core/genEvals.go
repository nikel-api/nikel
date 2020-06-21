// DO NOT TOUCH. This file is automatically generated.
package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

// getEvals queries via the evals endpoint
func getEvals(c *gin.Context) {
	offset := parseInt(c.Query("offset"), 0, math.MaxInt64, 0)
	limit := parseInt(c.Query("limit"), 1, TOPLIMIT, DEFAULTLIMIT)

	data := autoQuery(
		database.EvalsData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var evals []Eval
	json.Unmarshal(res, &evals)
	if len(evals) == 0 {
		sendEmptySuccess(c)
	} else {
		sendSuccess(c, evals)
	}
}

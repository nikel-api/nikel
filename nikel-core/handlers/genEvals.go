package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/database"
	"github.com/nikel-api/nikel/nikel-core/query"
	"github.com/nikel-api/nikel/nikel-core/response"
	"math"
)

// GetEvals queries via the evals endpoint
func GetEvals(c *gin.Context) {
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TopLimit, config.DefaultLimit)

	data := query.AutoQuery(
		database.DB.EvalsData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var evals []database.Eval
	json.Unmarshal(res, &evals)
	if len(evals) == 0 {
		response.SendEmptySuccess(c)
	} else {
		response.SendSuccess(c, evals)
	}
}

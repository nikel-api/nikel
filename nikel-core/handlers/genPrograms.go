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

// GetPrograms queries via the programs endpoint
func GetPrograms(c *gin.Context) {
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TopLimit, config.DefaultLimit)

	data := query.AutoQuery(
		database.DB.ProgramsData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)

	var programs []database.Program

	err := json.Unmarshal(res, &programs)
	if err != nil {
		response.SendError(c)
		return
	}

	if len(programs) == 0 {
		response.SendEmptySuccess(c)
	} else {
		response.SendSuccess(c, programs)
	}
}

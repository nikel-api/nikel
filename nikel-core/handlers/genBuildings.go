package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/database"
	"github.com/nikel-api/nikel/nikel-core/query"
	"github.com/nikel-api/nikel/nikel-core/response"
	"math"
)

// GetBuildings queries via the buildings endpoint
func GetBuildings(c *gin.Context) {
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TopLimit, config.DefaultLimit)

	data := query.AutoQuery(
		database.DB.BuildingsData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)

	var buildings []database.Building

	err := json.Unmarshal(res, &buildings)
	if err != nil {
		response.SendError(c)
		return
	}

	if len(buildings) == 0 {
		response.SendEmptySuccess(c)
	} else {
		response.SendSuccess(c, buildings)
	}
}

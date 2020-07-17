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

// GetServices queries via the services endpoint
func GetServices(c *gin.Context) {
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TopLimit, config.DefaultLimit)

	data := query.AutoQuery(
		database.DB.ServicesData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var services []database.Service
	json.Unmarshal(res, &services)
	if len(services) == 0 {
		response.SendEmptySuccess(c)
	} else {
		response.SendSuccess(c, services)
	}
}

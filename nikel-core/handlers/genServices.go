// This file is automatically generated using the codegen module.
// Please do not edit!

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
	// get offset and limit
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TopLimit, config.DefaultLimit)

	// query data
	data := query.AutoQuery(
		database.DB.ServicesData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	// convert data to bytes
	// this is expensive but reforming the data through an additional unmarshal is necessary
	res, _ := json.Marshal(data)

	// initialize payload variable
	var services []database.Service

	// unmarshal on a specific struct defined for this payload
	err := json.Unmarshal(res, &services)
	if err != nil {
		response.SendError(c)
		return
	}

	if len(services) == 0 {
		// send empty array
		response.SendEmptySuccess(c)
	} else {
		// send payload
		response.SendSuccess(c, services)
	}
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/database"
	"github.com/nikel-api/nikel/nikel-core/query"
	"github.com/nikel-api/nikel/nikel-core/response"
	"math"
)

// GetParking queries via the parking endpoint
func GetParking(c *gin.Context) {
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TopLimit, config.DefaultLimit)

	data := query.AutoQuery(
		database.DB.ParkingData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)

	var parking []database.Parking

	err := json.Unmarshal(res, &parking)
	if err != nil {
		response.SendError(c)
		return
	}

	if len(parking) == 0 {
		response.SendEmptySuccess(c)
	} else {
		response.SendSuccess(c, parking)
	}
}

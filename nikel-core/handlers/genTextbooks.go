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

// GetTextbooks queries via the textbooks endpoint
func GetTextbooks(c *gin.Context) {
	offset := query.ParseInt(c.Query("offset"), 0, math.MaxInt32, 0)
	limit := query.ParseInt(c.Query("limit"), 1, config.TOPLIMIT, config.DEFAULTLIMIT)

	data := query.AutoQuery(
		database.DB.TextbooksData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var textbooks []database.Textbook
	json.Unmarshal(res, &textbooks)
	if len(textbooks) == 0 {
		response.SendEmptySuccess(c)
	} else {
		response.SendSuccess(c, textbooks)
	}
}

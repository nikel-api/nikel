package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getStatus(c *gin.Context) {
	response := Metric{}
	err := urlToStruct(MetricApi, &response)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusInternalServerError,
			"status_message": "api metrics not found",
			"response": gin.H{
				"requests": nil,
			},
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response": gin.H{
			"requests": response.Value,
		},
	})
}

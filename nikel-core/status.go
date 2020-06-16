package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getStatus(c *gin.Context) {
	response := Metric{}
	err := urlToStruct(MetricApi, &response)
	if err != nil {
		sendServerError(c)
		return
	}
	sendSuccess(c, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response": gin.H{
			"requests": response.Value,
		},
	})
}

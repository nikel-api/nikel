package main

import (
	"github.com/gin-gonic/gin"
)

func getStatus(c *gin.Context) {
	response := Metric{}
	err := urlToStruct(MetricApi, &response)
	if err != nil {
		sendServerError(c)
		return
	}
	sendSuccess(c, gin.H{
		"requests": response.Value,
	})
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var HttpClient *http.Client
var MetricApi string

func init() {
	loadVals()
	HttpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	MetricApi = os.Getenv("METRICAPI")
}

func main() {
	router := gin.Default()
	router.Use(usageMetrics())
	router.GET("api/status", getStatus)
	router.GET("api/courses", getCourses)
	router.GET("api/courses/:p1", p1Handler)
	router.Run()
}

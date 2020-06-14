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
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://docs.nikel.ml")
	})
	router.GET("api/status", getStatus)
	router.GET("api/courses", getCourses)
	router.GET("api/courses/:p1", courseHandler)
	router.GET("api/buildings", getBuildings)
	router.GET("api/buildings/:p1", buildingHandler)
	router.GET("api/food", getFood)
	router.GET("api/food/:p1", foodHandler)
	router.GET("api/accessibility", getAccessibility)
	router.GET("api/accessibility/:p1", accessibilityHandler)
	router.Run()
}

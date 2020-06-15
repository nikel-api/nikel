package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
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
	// Handle rate limits
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  5,
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)
	middleware := mgin.NewMiddleware(instance)

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(middleware)
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
	router.GET("api/parking", getParking)
	router.GET("api/parking/:p1", parkingHandler)
	router.GET("api/accessibility", getAccessibility)
	router.GET("api/accessibility/:p1", accessibilityHandler)
	router.Run()
}

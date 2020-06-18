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
	ratelimit := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  10,
	}

	rateStore := memory.NewStore()
	rateInstance := limiter.New(rateStore, ratelimit)
	rateMiddleware := mgin.NewMiddleware(rateInstance)

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(rateMiddleware)
	router.Use(usageMetrics())
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://docs.nikel.ml")
	})

	router.GET("api/metrics", getMetrics)
	router.GET("api/courses", getCourses)
	router.GET("api/textbooks", getTextbooks)
	router.GET("api/buildings", getBuildings)
	router.GET("api/food", getFood)
	router.GET("api/parking", getParking)
	router.GET("api/accessibility", getAccessibility)
	router.GET("api/exams", getExams)
	router.Run()
}

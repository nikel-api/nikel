package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
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

	cacheStore := persistence.NewInMemoryStore(persistence.DEFAULT)

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(rateMiddleware)
	router.Use(usageMetrics())
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://docs.nikel.ml")
	})

	router.GET("api/status", getStatus)
	router.GET("api/metrics", getMetrics)
	router.GET("api/courses", cache.CachePageWithoutHeader(cacheStore, DEFAULTTTL, getCourses))
	router.GET("api/textbooks", cache.CachePageWithoutHeader(cacheStore, DEFAULTTTL, getTextbooks))
	router.GET("api/buildings", cache.CachePageWithoutHeader(cacheStore, DEFAULTTTL, getBuildings))
	router.GET("api/food", cache.CachePageWithoutHeader(cacheStore, DEFAULTTTL, getFood))
	router.GET("api/parking", cache.CachePageWithoutHeader(cacheStore, DEFAULTTTL, getParking))
	router.GET("api/accessibility", cache.CachePageWithoutHeader(cacheStore, DEFAULTTTL, getAccessibility))
	router.GET("api/exams", cache.CachePageWithoutHeader(cacheStore, DEFAULTTTL, getExams))
	router.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"time"
)

func init() {
	loadVals()
}

func main() {
	// set ratelimit at 20 req/s
	ratelimit := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  20,
	}

	rateStore := memory.NewStore()
	rateInstance := limiter.New(rateStore, ratelimit)
	rateMiddleware := mgin.NewMiddleware(rateInstance)

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(rateMiddleware)

	// define routes
	router.GET("api/metrics", getMetrics)
	router.GET("api/courses", getCourses)
	router.GET("api/textbooks", getTextbooks)
	router.GET("api/buildings", getBuildings)
	router.GET("api/food", getFood)
	router.GET("api/parking", getParking)
	router.GET("api/services", getServices)
	router.GET("api/exams", getExams)
	router.GET("api/evals", getEvals)

	// run server
	router.Run()
}

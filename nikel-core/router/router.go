package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/nikel-api/nikel/nikel-core/metrics"
	"github.com/nikel-api/nikel/nikel-core/response"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"os"
	"strconv"
	"time"
)

// GetRouter returns nikel-core's router
func GetRouter() *gin.Engine {
	router := gin.Default()

	// set ratelimit
	ratelimitValueStr := os.Getenv("RATELIMIT")
	if len(ratelimitValueStr) != 0 {
		ratelimitValueInt, err := strconv.Atoi(ratelimitValueStr)

		if err != nil {
			panic(err)
		}

		if ratelimitValueInt < 1 {
			panic(fmt.Errorf("nikel-core: invalid ratelimit value %d", ratelimitValueInt))
		}

		ratelimit := limiter.Rate{
			Period: 1 * time.Second,
			Limit:  int64(ratelimitValueInt),
		}

		rateStore := memory.NewStore()
		rateInstance := limiter.New(rateStore, ratelimit)
		rateMiddleware := mgin.NewMiddleware(rateInstance)
		router.ForwardedByClientIP = true
		router.Use(rateMiddleware)
	}

	// define routes
	router.GET("api/metrics", metrics.GetMetrics)
	router.GET("api/courses", handlers.GetCourses)
	router.GET("api/textbooks", handlers.GetTextbooks)
	router.GET("api/buildings", handlers.GetBuildings)
	router.GET("api/food", handlers.GetFood)
	router.GET("api/parking", handlers.GetParking)
	router.GET("api/services", handlers.GetServices)
	router.GET("api/exams", handlers.GetExams)
	router.GET("api/evals", handlers.GetEvals)
	router.NoRoute(response.SendNotFound)

	return router
}

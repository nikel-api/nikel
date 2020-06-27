package main

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/nikel-api/nikel/nikel-core/metrics"
	"github.com/nikel-api/nikel/nikel-core/response"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"os"
	"time"
)

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

	// new relic apm monitoring
	newRelicLicense := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if len(newRelicLicense) != 0 {
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName("Nikel API"),
			newrelic.ConfigLicense(newRelicLicense),
		)
		if err != nil {
			panic(err)
		}
		router.Use(nrgin.Middleware(app))
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

	// run server
	err := router.Run()
	if err != nil {
		panic(err)
	}
}

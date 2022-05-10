package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/database"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/nikel-api/nikel/nikel-core/metrics"
	"github.com/nikel-api/nikel/nikel-core/response"
)

// SetRoutes sets the routes
func (r *Router) SetRoutes() *Router {
	// dispatch queryable routes to cached engine
	r.Cached.GET("api/courses", func(c *gin.Context) {
		handlers.Get[database.Course](c, database.DB.CoursesData)
	})
	r.Cached.GET("api/programs", func(c *gin.Context) {
		handlers.Get[database.Program](c, database.DB.ProgramsData)
	})
	r.Cached.GET("api/textbooks", func(c *gin.Context) {
		handlers.Get[database.Textbook](c, database.DB.TextbooksData)
	})
	r.Cached.GET("api/buildings", func(c *gin.Context) {
		handlers.Get[database.Building](c, database.DB.BuildingsData)
	})
	r.Cached.GET("api/food", func(c *gin.Context) {
		handlers.Get[database.Food](c, database.DB.FoodData)
	})
	r.Cached.GET("api/parking", func(c *gin.Context) {
		handlers.Get[database.Parking](c, database.DB.ParkingData)
	})
	r.Cached.GET("api/services", func(c *gin.Context) {
		handlers.Get[database.Service](c, database.DB.ServicesData)
	})
	r.Cached.GET("api/exams", func(c *gin.Context) {
		handlers.Get[database.Exam](c, database.DB.ExamsData)
	})
	r.Cached.GET("api/evals", func(c *gin.Context) {
		handlers.Get[database.Eval](c, database.DB.EvalsData)
	})

	// support favicon
	r.Cached.StaticFile(
		"favicon.ico",
		config.PathPrefix+config.FaviconPath,
	)

	// dispatch non-queryable routes to uncached engine
	r.Uncached.GET("api/metrics", metrics.GetMetrics)

	// send 404 NotFound for no routes
	r.Engine.NoRoute(response.SendNotFound)

	fmt.Println("[NIKEL-CORE] All routes initialized.")
	return r
}

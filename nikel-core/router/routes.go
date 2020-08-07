package router

import (
	"fmt"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/nikel-api/nikel/nikel-core/metrics"
	"github.com/nikel-api/nikel/nikel-core/response"
)

// SetRoutes sets the routes
func (r *Router) SetRoutes() *Router {
	// dispatch queryable routes to cached engine
	r.Cached.GET("api/courses", handlers.GetCourses)
	r.Cached.GET("api/programs", handlers.GetPrograms)
	r.Cached.GET("api/textbooks", handlers.GetTextbooks)
	r.Cached.GET("api/buildings", handlers.GetBuildings)
	r.Cached.GET("api/food", handlers.GetFood)
	r.Cached.GET("api/parking", handlers.GetParking)
	r.Cached.GET("api/services", handlers.GetServices)
	r.Cached.GET("api/exams", handlers.GetExams)
	r.Cached.GET("api/evals", handlers.GetEvals)

	// dispatch non-queryable routes to uncached engine
	r.Uncached.GET("api/metrics", metrics.GetMetrics)

	// send 404 NotFound for no routes
	r.Engine.NoRoute(response.SendNotFound)

	fmt.Println("[NIKEL-CORE] All routes initialized.")
	return r
}

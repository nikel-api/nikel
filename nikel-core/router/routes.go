package router

import (
	"fmt"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/nikel-api/nikel/nikel-core/metrics"
	"github.com/nikel-api/nikel/nikel-core/response"
)

// SetRoutes sets the routes
func (r *Router) SetRoutes() *Router {
	r.Cached.GET("api/courses", handlers.GetCourses)
	r.Cached.GET("api/textbooks", handlers.GetTextbooks)
	r.Cached.GET("api/buildings", handlers.GetBuildings)
	r.Cached.GET("api/food", handlers.GetFood)
	r.Cached.GET("api/parking", handlers.GetParking)
	r.Cached.GET("api/services", handlers.GetServices)
	r.Cached.GET("api/exams", handlers.GetExams)
	r.Cached.GET("api/evals", handlers.GetEvals)
	r.Uncached.GET("api/metrics", metrics.GetMetrics)
	r.Engine.NoRoute(response.SendNotFound)

	fmt.Println("[NIKEL-CORE] All routes initialized.")
	return r
}

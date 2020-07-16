package router

import (
	"fmt"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/nikel-api/nikel/nikel-core/metrics"
	"github.com/nikel-api/nikel/nikel-core/response"
)

// SetRoutes sets the routes
func (r *Router) SetRoutes() *Router {
	r.Engine.GET("api/metrics", metrics.GetMetrics)
	r.Engine.GET("api/courses", handlers.GetCourses)
	r.Engine.GET("api/textbooks", handlers.GetTextbooks)
	r.Engine.GET("api/buildings", handlers.GetBuildings)
	r.Engine.GET("api/food", handlers.GetFood)
	r.Engine.GET("api/parking", handlers.GetParking)
	r.Engine.GET("api/services", handlers.GetServices)
	r.Engine.GET("api/exams", handlers.GetExams)
	r.Engine.GET("api/evals", handlers.GetEvals)
	r.Engine.NoRoute(response.SendNotFound)

	fmt.Println("[NIKEL-CORE] All routes initialized.")
	return r
}

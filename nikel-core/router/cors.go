package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
)

// SetAllowCors allows all origins
func (r *Router) SetAllowCors() *Router {
	corsMiddleware := cors.Default()

	// Attach both cached and uncached groups
	r.Cached.Use(corsMiddleware)
	r.Uncached.Use(corsMiddleware)

	fmt.Println("[NIKEL-CORE] Set to allow all origins.")
	return r
}

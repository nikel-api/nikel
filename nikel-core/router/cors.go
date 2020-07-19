package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
)

// SetAllowCors allows all origins
func (r *Router) SetAllowCors() *Router {
	r.Engine.Use(cors.Default())
	fmt.Println("[NIKEL-CORE] Set to allow all origins.")
	return r
}

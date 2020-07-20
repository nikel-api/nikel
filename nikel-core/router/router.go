package router

import (
	"github.com/gin-gonic/gin"
)

// Router contains bindings for gin's router engine
type Router struct {
	Engine   *gin.Engine
	Cached   *gin.RouterGroup
	Uncached *gin.RouterGroup
}

// NewRouter returns nikel-core's base router
func NewRouter() *Router {
	r := gin.Default()
	cached := r.Group("/")
	uncached := r.Group("/")
	return &Router{
		Engine:   r,
		Cached:   cached,
		Uncached: uncached,
	}
}

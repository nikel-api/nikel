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

	// there are two groups defined
	// cached has the cache middleware attached
	// uncached does not
	return &Router{
		Engine:   r,
		Cached:   r.Group("/"),
		Uncached: r.Group("/"),
	}
}

package router

import (
	"github.com/gin-gonic/gin"
)

// Router contains bindings for gin's router engine
type Router struct {
	Engine *gin.Engine
}

// NewRouter returns nikel-core's base router
func NewRouter() *Router {
	return &Router{
		Engine: gin.Default(),
	}
}

package router

import (
	"fmt"
	"github.com/gin-contrib/gzip"
)

// SetGzip gzips all api responses
func (r *Router) SetGzip() *Router {
	// Attach both cached and uncached groups
	r.Cached.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Uncached.Use(gzip.Gzip(gzip.DefaultCompression))

	fmt.Println("[NIKEL-CORE] Set to gzip all api responses.")
	return r
}

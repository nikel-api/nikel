package main

import (
	"github.com/nikel-api/nikel/nikel-core/router"
)

func main() {
	// get router
	r := router.
		NewRouter().
		SetRateLimiter().
		SetLevelDBCache().
		SetAllowCors().
		SetRoutes()

	// run server
	err := r.Engine.Run()
	if err != nil {
		panic(err)
	}
}

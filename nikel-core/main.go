package main

import (
	"github.com/nikel-api/nikel/nikel-core/router"
)

func main() {
	// get router
	r := router.
		NewRouter().
		SetAllowCors().
		SetRateLimiter().
		SetRoutes()

	// run server
	err := r.Engine.Run()
	if err != nil {
		panic(err)
	}
}

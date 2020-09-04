package main

import (
	"github.com/nikel-api/nikel/nikel-core/router"
)

func main() {
	// get router and attach middlewares
	r := router.
		NewRouter().
		SetRateLimiter().
		SetLevelDBCache().
		SetAllowCors().
		SetGZip().
		SetRoutes()

	// run server
	err := r.Engine.Run()
	if err != nil {
		panic(err)
	}
}

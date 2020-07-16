package main

import (
	"github.com/nikel-api/nikel/nikel-core/router"
)

func main() {
	// get router
	r := router.GetRouter()

	// run server
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

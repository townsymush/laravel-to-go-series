package main

import (
	"log"
	"net/http"

	"github.com/townsymush/laravel-to-go-series/part-2-routing-and-handlers/app/controllers"
)

func main() {
	// set up a route
	http.HandleFunc("/", controllers.HomeHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

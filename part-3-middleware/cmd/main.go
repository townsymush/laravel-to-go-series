package main

import (
	"log"
	"net/http"

	"github.com/townsymush/laravel-to-go-series/part-3-middleware/app/controllers"
	"github.com/townsymush/laravel-to-go-series/part-3-middleware/app/middleware"
)

func main() {

	s := http.NewServeMux()

	s.HandleFunc("/", controllers.HomeHandler)

	authWrapper := middleware.NewAuth(s)

	log.Fatal(http.ListenAndServe(":7777", authWrapper))
}

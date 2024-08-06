package main

import (
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app"
	"log"
	"net/http"
)

func main() {
	a := app.New()

	err := http.ListenAndServe(":7777", a.Router())
	if err != nil {
		log.Fatal(err) //todo improve this
	}
}

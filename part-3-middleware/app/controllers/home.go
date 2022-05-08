package controllers

import (
	"log"
	"net/http"

	"github.com/townsymush/laravel-to-go-series/part-3-middleware/app/response"
)

// Home Handler is the handler for our `/` url
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	m := "Hello World - Welcome to my first Go homepage"

	// Create response from message
	resp := response.NewJSON(http.StatusOK, m)

	if err := resp.WriteResponse(w); err != nil {
		log.Fatalf("failed to write a response: %s", err.Error())
	}
}

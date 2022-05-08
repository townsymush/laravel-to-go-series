package controllers

import (
	"log"
	"net/http"
)

// Home Handler is the handler for our `/` url
func HomeHandler(w http.ResponseWriter, _ *http.Request) {

	m := "Hello World - Welcome to my first Go homepage"

	// Create response from message
	resp := NewJSONResponse(http.StatusOK, m)

	if err := resp.WriteResponse(w); err != nil {
		log.Fatalf("failed to write a response: %s", err.Error())
	}
}

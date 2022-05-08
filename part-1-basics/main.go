package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	home := func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "Hello World - Welcome to my first Go homepage")
	}
	// set up a route
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

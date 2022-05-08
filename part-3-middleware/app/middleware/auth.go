package middleware

import (
	"fmt"
	"net/http"

	"github.com/townsymush/laravel-to-go-series/part-3-middleware/app/response"
)

// header constants
const (
	HeaderAuthUsername = "x-credentials-username"
	HeaderAuthPassword = "x-credentials-password"
)

// Auth is a middleware handler which verifies a username
// and password headers.
type Auth struct {
	handler http.Handler
}

// NewAuth accepts a handler and decorates it with Authentication header check middleware
func NewAuth(decorates http.Handler) *Auth {
	return &Auth{handler: decorates}
}

// ServeHTTP is used to call the handler chain once authentication headers
// have been checked
func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := r.Header.Get(HeaderAuthUsername)
	p := r.Header.Get(HeaderAuthPassword)

	if u == "my-username" && p == "my-password" {
		a.handler.ServeHTTP(w, r)
		return
	}

	// write the response
	resp := response.NewJSON(http.StatusUnauthorized, "failed to authenticate")
	err := resp.WriteResponse(w)
	if err != nil {
		fmt.Println("failed to write json response in auth middleware")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

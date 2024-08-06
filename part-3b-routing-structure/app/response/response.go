package response

import (
	"net/http"
)

// Responder is the expected behaviour from a http response struct
type Responder interface {
	WriteResponse(w http.ResponseWriter) error
}

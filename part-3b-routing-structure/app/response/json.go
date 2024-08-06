package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponse holds our json payload for marshalling
type JSONResponse struct {
	statusCode int
	Message    string `json:"Message"`
}

// NewJSON is a constructor for a json Response
func NewJSON(statusCode int, message string) JSONResponse {
	return JSONResponse{
		statusCode: statusCode,
		Message:    message,
	}
}

// WriteResponse writes the structs payload to the Response
func (j *JSONResponse) WriteResponse(w http.ResponseWriter) error {
	// Sets json header
	w.Header().Add("Content-Type", "application/json")

	// writes the jsonResponse struct into json bytes
	b, err := json.Marshal(j)
	if err != nil {
		return fmt.Errorf("could not marshal jsonResponse %w", err)
	}

	// Write status code for response
	w.WriteHeader(j.statusCode)

	// write the payload to the response
	_, err = w.Write(b)
	if err != nil {
		return fmt.Errorf("could not write jsonResponse %w", err)
	}

	return nil
}

package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// JSONError represents a list of errors for the client
type JSONError struct {
	Errors []string `json:"errors"`
}

// NewJSONError returns a JSONError instantiated with the given strings as
// errors.
func NewJSONError(errs ...string) JSONError {
	return JSONError{Errors: errs}
}

func respond(w http.ResponseWriter, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		respond(w, http.StatusInternalServerError, NewJSONError("Internal server error"))
		return
	}
	w.WriteHeader(status)
	if data == nil {
		return
	}
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("respond: ", err)
	}
}

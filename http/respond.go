package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("respond: ", err)
	}
}

func respondOK(w http.ResponseWriter, data interface{}) {
	respond(w, http.StatusOK, data)
}

func respondBadRequest(w http.ResponseWriter) {
	respond(w, http.StatusBadRequest, NewJSONError("Bad request"))
}

func respondUnauthorized(w http.ResponseWriter) {
	respond(w, http.StatusUnauthorized, NewJSONError("Unauthorized"))
}

func respondInternalServerError(w http.ResponseWriter) {
	respond(w, http.StatusInternalServerError, NewJSONError("Internal server error"))
}

func respondValidationErrors(w http.ResponseWriter, err error) {
	errs, ok := err.(govalidator.Errors)
	if !ok {
		respondInternalServerError(w)
		return
	}
	errList := errs.Errors()
	errStrs := make([]string, len(errs))
	for i := range errList {
		errStrs[i] = errList[i].Error()
	}
	respond(w, http.StatusBadRequest, NewJSONError(errStrs...))
}

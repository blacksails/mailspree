package http

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/blacksails/mailspree"
)

func (s server) mailHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse incoming request
		var m mailspree.Message
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			respondBadRequest(w)
			return
		}

		// Validate request
		if valid, err := govalidator.ValidateStruct(m); !valid {
			respondValidationErrors(w, err)
			return
		}

		// Send the email
		if err := s.mailingProvider.SendEmail(m); err != nil {
			// If it fails tell the client
			respond(w, http.StatusServiceUnavailable, newJSONError(err.Error()))
			return
		}
		respondOK(w, nil)
	})
}

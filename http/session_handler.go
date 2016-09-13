package http

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/blacksails/mailspree"
)

func (s server) sessionHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse incoming request
		var ns mailspree.NewSession
		if err := json.NewDecoder(r.Body).Decode(&ns); err != nil {
			respondBadRequest(w)
			return
		}

		// Validation
		if valid, err := govalidator.ValidateStruct(ns); !valid {
			respondValidationErrors(w, err)
			return
		}

		// Lookup user
		user, err := s.userService.Find(ns.Username)
		if err != nil {
			respondUnauthorized(w)
			return
		}

		// Authenticate user
		token, err := s.authService.Authenticate(user, ns.Password)
		if err != nil {
			respondUnauthorized(w)
			return
		}

		// Create session and send it
		ses := mailspree.Session{Username: user.Username, Token: token}
		respondOK(w, ses)
	})
}

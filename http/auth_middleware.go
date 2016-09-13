package http

import (
	"errors"
	"net/http"
	"strings"
)

func (s server) ensureAuth() adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr, err := getBearerToken(r)
			if err != nil {
				respondUnauthorized(w)
				return
			}
			_, err = s.authService.Validate(tokenStr, s.userService)
			if err != nil {
				respondUnauthorized(w)
				return
			}
			// TODO: maybe set the user in the context so that we can easily
			// get the current user downstream.
			h.ServeHTTP(w, r)
		})
	}
}

func getBearerToken(r *http.Request) (string, error) {
	h := r.Header.Get("Authorization")
	if h == "" {
		return "", errors.New("No authorization header was found")
	}
	hParts := strings.Split(h, " ")
	if len(hParts) != 2 || strings.ToLower(hParts[0]) != "bearer" {
		return "", errors.New("Authorization header method must be Bearer <token>")
	}
	return hParts[1], nil
}

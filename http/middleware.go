package http

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	rscors "github.com/rs/cors"
)

type adapter func(http.Handler) http.Handler

func adapt(h http.Handler, adapters ...adapter) http.Handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}
	return h
}

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

func cors() adapter {
	return rscors.New(rscors.Options{
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization"},
	}).Handler
}

func logging() adapter {
	return func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
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

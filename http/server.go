package http

import (
	"net/http"

	"github.com/blacksails/mailspree"
	"github.com/gorilla/mux"
)

// Server is a http.Handler which implements the mailspree service
type server struct {
	mailingProvider mailspree.MailingProvider
	userService     mailspree.UserService
	authService     mailspree.AuthService
}

// NewServer returns a server configured with the given mailing
// provider.
func NewServer(mp mailspree.MailingProvider, us mailspree.UserService, as mailspree.AuthService) http.Handler {
	return server{mailingProvider: mp, userService: us, authService: as}
}

// ServeHTTP is the Server http.Handler implementation
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.Handle("/send", adapt(s.mailHandler(), s.ensureAuth())).Methods("POST")
	router.Handle("/session", s.sessionHandler()).Methods("POST")
	handler := adapt(router, jsonContentType())
	handler.ServeHTTP(w, r)
}

// ListenAndServe simply wraps the net/http ListenAndServe
func ListenAndServe(addr string, h http.Handler) error {
	return http.ListenAndServe(addr, h)
}

type adapter func(http.Handler) http.Handler

func adapt(h http.Handler, adapters ...adapter) http.Handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}
	return h
}

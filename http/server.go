package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blacksails/mailspree"
	"github.com/gorilla/mux"
)

// Server is a http.Handler which implements the mailspree service
type Server struct {
	mailingProviders []mailspree.MailingProvider
}

func (s Server) mailHandler(w http.ResponseWriter, r *http.Request) {
	var e mailspree.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// ServeHTTP is the Server http.Handler implementation
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := mux.NewRouter()
	mux.HandleFunc("/send", s.mailHandler).Methods("POST")
	mux.ServeHTTP(w, r)
}

// ListenAndServe simply wraps the net/http ListenAndServe
func ListenAndServe(addr string, h http.Handler) error {
	return http.ListenAndServe(addr, h)
}

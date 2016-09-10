package http

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/blacksails/mailspree"
	"github.com/gorilla/mux"
)

// Server is a http.Handler which implements the mailspree service
type server struct {
	mailingProvider mailspree.MailingProvider
}

// NewServer returns a server configured with the given mailing
// provider.
func NewServer(mp mailspree.MailingProvider) http.Handler {
	return server{mailingProvider: mp}
}

func (s server) mailHandler(w http.ResponseWriter, r *http.Request) {
	var m mailspree.Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		respond(w, http.StatusBadRequest, NewJSONError("Bad request"))
		return
	}

	valid, err := govalidator.ValidateStruct(m)
	if !valid {
		errs := err.(govalidator.Errors).Errors()
		errStrs := make([]string, len(errs))
		for i := range errs {
			errStrs[i] = errs[i].Error()
		}
		respond(w, http.StatusBadRequest, NewJSONError(errStrs...))
		return
	}

	// Send the email
	err = s.mailingProvider.SendEmail(m)
	if err != nil {
		// If it fails tell the client
		respond(w, http.StatusServiceUnavailable, NewJSONError(err.Error()))
		return
	}
	// If it succeeds respond OK
	respond(w, http.StatusOK, nil)
}

// ServeHTTP is the Server http.Handler implementation
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/send", s.mailHandler).Methods("POST")
	handler := jsonContentType(router)
	handler.ServeHTTP(w, r)
}

// ListenAndServe simply wraps the net/http ListenAndServe
func ListenAndServe(addr string, h http.Handler) error {
	return http.ListenAndServe(addr, h)
}

func jsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

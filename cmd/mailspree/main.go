// The command mailspree starts a mailspree server. It is configured with the
// following required env variables. See the mailspree-vars.env.example file
// in the root directory for an example.
//
// Mailgun configuration
//
// The two following environement variables needs to be set, so that the
// mailgun provider knows how to send mails. Both are required to start the
// server.
//
//	MAILGUN_DOMAIN
//	MAILGUN_APIKEY
//
// Sendgrid configuration
//
// The next environement variables needs to be set, so that the
// sendgrid provider knows how to send mails. It is required.
//
//	SENDGRID_APIKEY
//
// Mailspree configuration
//
// The last three environment variables we need are the following.
//
//	MAILSPREE_USER
//	MAILSPREE_PASS
//	MAILSPREE_PK
//
// The first two are the username and password, used to login to the service.
// The last one is the private key used to sign authentication tokens. All
// three are required.
//
// Provider priority
//
// The environment variable PROVIDER_PRIORITY decides what mailing provider
// gets first priority. In order to make mailspree use mailgun as the primary
// provider, set the variable like this.
//
//	PROVIDER_PRIORITY=mg
//
// or
//
//	PROVIDER_PRIORITY=mailgun
//
// To make sendgrid the primary provider set it to 'sg' or 'sendgrid'. If it is
// unset, mailspree defaults to using mailgun.
package main

import (
	"log"
	"os"
	"strings"

	"github.com/blacksails/mailspree"
	"github.com/blacksails/mailspree/http"
	"github.com/blacksails/mailspree/jwt"
	"github.com/blacksails/mailspree/mailgun"
	"github.com/blacksails/mailspree/sendgrid"
)

func main() {

	var mps mailspree.MailingProviders

	// Mailgun configuration
	mgDomain := os.Getenv("MAILGUN_DOMAIN")
	mgKey := os.Getenv("MAILGUN_APIKEY")
	if mgDomain == "" || mgKey == "" {
		log.Fatalln("Please provide a mailgun domain and api key")
	}
	mgProvider := mailgun.MailingProvider{Domain: mgDomain, APIKey: mgKey}

	// Sendgrid configuration
	sgKey := os.Getenv("SENDGRID_APIKEY")
	if sgKey == "" {
		log.Fatalln("Please provide a sendgrid api key")
	}
	sgProvider := sendgrid.MailingProvider{APIKey: sgKey}

	// Priority configuration
	prio := os.Getenv("PROVIDER_PRIORITY")

	switch strings.ToLower(prio) {
	case "sendgrid", "sg":
		mps = append(mps, sgProvider, mgProvider)
	case "mailgrid", "mg":
		mps = append(mps, mgProvider, sgProvider)
	default:
		mps = append(mps, mgProvider, sgProvider)
	}

	// Wrap providers in circuit breakers
	mpsWithCBs := make(mailspree.MailingProviders, len(mps))
	for i, mp := range mps {
		mpsWithCBs[i] = mailspree.NewCircuitBreaker(mp, mailspree.NewCircuitBreakerTimer())
	}

	// Mailspree user
	msUser := os.Getenv("MAILSPREE_USER")
	msPass := os.Getenv("MAILSPREE_PASS")
	if msUser == "" || msPass == "" {
		log.Fatalln("Please provide a user and password")
	}
	us := mailspree.SimpleUserService{User: mailspree.NewUser(msUser, msPass)}

	// Mailspree authentication service
	msPK := os.Getenv("MAILSPREE_PK")
	if msPK == "" {
		log.Fatalln("Please provide a secret key for signing tokens")
	}
	as := jwt.AuthService{PrivateKey: msPK}

	err := http.ListenAndServe(":8080", http.NewServer(mpsWithCBs, us, as))
	if err != nil {
		log.Fatalln(err)
	}
}

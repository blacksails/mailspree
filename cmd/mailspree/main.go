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

	err := http.ListenAndServe(":8080", http.NewServer(mps, us, as))
	if err != nil {
		log.Fatalln(err)
	}
}

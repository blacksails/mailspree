package main

import (
	"log"
	"os"

	"github.com/blacksails/mailspree"
	"github.com/blacksails/mailspree/http"
	"github.com/blacksails/mailspree/mailgun"
)

func main() {

	mps := mailspree.MailingProviders{}

	// Mailgun configuration
	d := os.Getenv("MAILGUN_DOMAIN")
	key := os.Getenv("MAILGUN_APIKEY")
	switch {
	case d != "" && key != "":
		mps = append(mps, mailgun.MailingProvider{
			Domain: d,
			APIKey: key,
		})
	case d == "" && key != "", d != "" && key == "":
		log.Fatalln("When using the mailgun service, please set MAILGUN_DOMAIN and MAILGUN_APIKEY")
	}

	if len(mps) == 0 {
		log.Fatalln("We need at least one mailing provider in order to run mailspree")
	}

	err := http.ListenAndServe(":8080", http.NewServer(mps))
	if err != nil {
		log.Fatalln(err)
	}
}

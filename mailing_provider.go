package mailspree

import (
	"errors"
	"log"
)

// MailingProvider abstacts different mailing providers.
type MailingProvider interface {
	SendEmail(Message) error
}

// MailingProviders is a list of MailingProvider implementations
type MailingProviders []MailingProvider

// SendEmail runs through the mailing providers and tries to send the email.
// Returns on first successful try. This also makes MailingProviders be a
// MailingProvider, which is a bit fun.
func (mps MailingProviders) SendEmail(m Message) error {
	for _, mp := range mps {
		err := mp.SendEmail(m)
		if err != nil {
			// TODO: make this a debug log entry
			log.Printf("DEBUG: mailing provider failed\nerror: %v", err)
		} else {
			return nil
		}
	}
	return errors.New("all mailing providers are failing")
}

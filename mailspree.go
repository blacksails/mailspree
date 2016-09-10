package mailspree

import (
	"bytes"
	"errors"
	"log"
)

// Message represents an email with all the related information.
type Message struct {
	From    Email   `json:"from"`
	To      []Email `json:"to"`
	Subject string  `json:"subject" valid:"required"`
	Body    string  `json:"body" valid:"required"`
}

// Email is simply a real name together with the address
type Email struct {
	Name    string `json:"name"`
	Address string `json:"address" valid:"email,required"`
}

// String returns a string representation of an Email
func (e Email) String() string {
	b := bytes.NewBufferString(e.Name)
	if e.Name != "" {
		b.WriteString(" <")
		b.WriteString(e.Address)
		b.WriteString(">")
	} else {
		b.WriteString(e.Address)
	}
	return b.String()
}

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

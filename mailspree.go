package mailspree

import "errors"

// Email represents an email with all the related information.
type Email struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// MailingProvider abstacts different mailing providers.
type MailingProvider interface {
	SendEmail(Email) error
}

// MailingProviders is a list of MailingProvider implementations
type MailingProviders []MailingProvider

// SendEmail runs through the mailing providers and tries to send the email.
// Returns on first successful try. This also makes MailingProviders be a
// MailingProvider, which is a bit fun.
func (mps MailingProviders) SendEmail(e Email) error {
	for _, mp := range mps {
		err := mp.SendEmail(e)
		if err == nil {
			return nil
		}
	}
	return errors.New("all mailing providers are failing")
}

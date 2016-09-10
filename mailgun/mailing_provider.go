package mailgun

import (
	"github.com/blacksails/mailspree"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

// MailingProvider is the mailgun implementation of mailspree.MailingProvider
type MailingProvider struct {
	Domain string
	APIKey string
}

// SendEmail uses the mailgun go api to send a mailspree.Message
func (mp MailingProvider) SendEmail(m mailspree.Message) error {
	mg := mailgun.NewMailgun(mp.Domain, mp.APIKey, "")
	msg := mg.NewMessage(m.From.String(), m.Subject, m.Body)
	for _, e := range m.To {
		msg.AddRecipient(e.String())
	}
	_, _, err := mg.Send(msg)
	return err
}

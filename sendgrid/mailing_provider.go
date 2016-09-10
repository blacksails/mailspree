package sendgrid

import (
	"errors"
	"log"

	"github.com/blacksails/mailspree"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// MailingProvider is a send grid mailing provider
type MailingProvider struct {
	APIKey string
}

// SendEmail uses the sendgrid API to send the given mailspree.Message
func (mp MailingProvider) SendEmail(m mailspree.Message) error {
	from := mail.NewEmail(m.From.Name, m.From.Address)
	content := mail.NewContent("text/plain", m.Body)
	p := mail.NewPersonalization()
	for _, e := range m.To {
		p.AddTos(mail.NewEmail(e.Name, e.Address))
	}

	sgm := mail.NewV3Mail()
	sgm.SetFrom(from)
	sgm.AddPersonalizations(p)
	sgm.Subject = m.Subject
	sgm.AddContent(content)

	req := sendgrid.GetRequest(mp.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	req.Method = "POST"
	req.Body = mail.GetRequestBody(sgm)
	_, err := sendgrid.API(req)
	if err != nil {
		msg := "Sendgrid mailing provider failed"
		log.Printf(msg)
		return errors.New(msg)
	}
	return nil
}

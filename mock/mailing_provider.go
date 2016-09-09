package mock

import (
	"errors"

	"github.com/blacksails/mailspree"
)

// MailingProvider is a mailing provider mock. It stores sent mail in a slice
// so that we can check if it is called. It can be set to fail by setting the
// Fail field to true.
type MailingProvider struct {
	SentEmail []mailspree.Email
	Fail      bool
}

// SendEmail adds the email to a slice for later inspection. If the Fail field
// is set, the email is not added, and we just fail.
func (mp *MailingProvider) SendEmail(e mailspree.Email) error {
	if mp.Fail {
		return errors.New("The mailing provider failed sending the email")
	}
	mp.SentEmail = append(mp.SentEmail, e)
	return nil
}

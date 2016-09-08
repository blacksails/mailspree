package mock

import (
	"errors"

	"github.com/blacksails/mailspree"
)

// SucceedingMailingProvider is a mock implementation of
// mailspree.MailingProvider, which simply appends sent email to a slice.
type SucceedingMailingProvider struct {
	SentEmail []mailspree.Email
}

// SendEmail is the function that SucceedingMailingProvider implements to
// satisfy the mailspree.MailingProvider interface.
func (mp *SucceedingMailingProvider) SendEmail(e mailspree.Email) error {
	mp.SentEmail = append(mp.SentEmail, e)
	return nil
}

// FailingMailingProvider is a mock implementation of
// mailspree.MailingProvider, which always returns an error.
type FailingMailingProvider struct{}

// SendEmail is the function that FailingMailingProvider implements to
// satisfy the mailspree.MailingProvider interface.
func (mp FailingMailingProvider) SendEmail(e mailspree.Email) error {
	return errors.New("This mailing provider failed to send the email")
}

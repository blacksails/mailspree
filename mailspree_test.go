package mailspree_test

import (
	"testing"

	"github.com/blacksails/mailspree"
	"github.com/blacksails/mailspree/mock"
)

func TestMailingProviders(t *testing.T) {
	// Positive case
	mp1 := mock.FailingMailingProvider{}
	mp2 := &mock.SucceedingMailingProvider{}
	mp3 := &mock.SucceedingMailingProvider{}
	mps := mailspree.MailingProviders{mp1, mp2, mp3}
	err := mps.SendEmail(mailspree.Email{})
	if err != nil {
		t.Error("the second provider should cause err to eq nil")
	}
	if len(mp2.SentEmail) != 1 {
		t.Error("the second provider should send the email")
	}
	if len(mp3.SentEmail) != 0 {
		t.Error("the third provider should not send the email")
	}

	// Negative case
	mps = mailspree.MailingProviders{mp1, mp1, mp1}
	err = mps.SendEmail(mailspree.Email{})
	if err == nil {
		t.Error("if all mailing providers fail we should return an error")
	}
}

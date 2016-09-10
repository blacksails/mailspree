package mock_test

import (
	"testing"

	"github.com/blacksails/mailspree"
	"github.com/blacksails/mailspree/mock"
)

func TestMailingProvider(t *testing.T) {
	mp := mock.MailingProvider{}
	e := mailspree.Message{}
	err := mp.SendEmail(e)
	if err != nil {
		t.Error("the mock should be configured to never fail")
	}
	if len(mp.SentEmail) != 1 {
		t.Error("there should now be one sent email in the mock")
	}
}

func TestMailingProviderFailure(t *testing.T) {
	mp := mock.MailingProvider{}
	e := mailspree.Message{}
	mp.Fail = true
	err := mp.SendEmail(e)
	if err == nil {
		t.Error("the mock should be configured to always fail")
	}
	if len(mp.SentEmail) != 0 {
		t.Error("there should not have been sent any email")
	}
}

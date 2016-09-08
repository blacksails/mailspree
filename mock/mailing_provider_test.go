package mock_test

import (
	"testing"

	"github.com/blacksails/mailspree"
	"github.com/blacksails/mailspree/mock"
)

func TestSucceedingMailingProvider(t *testing.T) {
	mp := mock.SucceedingMailingProvider{}
	e := mailspree.Email{}
	err := mp.SendEmail(e)
	if err != nil {
		t.Error("this mailing provider should always succeed")
	}
	if len(mp.SentEmail) != 1 {
		t.Error("after sending one mail the length of the SentEmail slice should be 1")
	}
}

func TestFailingMailingProvider(t *testing.T) {
	mp := mock.FailingMailingProvider{}
	e := mailspree.Email{}
	err := mp.SendEmail(e)
	if err == nil {
		t.Error("this mailing provider should always fail")
	}
}

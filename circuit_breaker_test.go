package mailspree_test

import (
	"testing"
	"time"

	"github.com/blacksails/mailspree"
	"github.com/blacksails/mailspree/mock"
)

func TestCircuitBreaker(t *testing.T) {
	// This test needs to sleep to avoid race conditions, so we skip it in
	// short test runs.
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	mp := mock.MailingProvider{}
	timer := mock.NewCircuitBreakerTimer()
	cb := mailspree.NewCircuitBreaker(&mp, timer)
	m := mailspree.Message{}

	// cb.state = closed
	for i := 0; i < 5; i++ {
		cb.SendEmail(m)
	}
	if len(mp.SentEmail) != 5 {
		t.Error("in closed state the emails should just be passed on to the mailing provider")
	}

	testThatWeAreInClosed(t, cb, &mp, timer, 0)
	mp.Fail = true
	cb.SendEmail(m)
	testThatWeAreInClosed(t, cb, &mp, timer, 1)
	mp.Fail = false
	cb.SendEmail(m)
	testThatWeAreInClosed(t, cb, &mp, timer, 0)
	mp.Fail = true
	cb.SendEmail(m)
	testThatWeAreInClosed(t, cb, &mp, timer, 1)
	cb.SendEmail(m)
	testThatWeAreInClosed(t, cb, &mp, timer, 2)
	cb.SendEmail(m)
	testThatWeAreInOpen(t, cb, &mp)
	timer.EndTimer()
	time.Sleep(time.Millisecond * 200)
	testThatWeAreInHalfOpen(t, cb, &mp)
}

func testThatWeAreInClosed(t *testing.T, cb mailspree.MailingProvider, mp *mock.MailingProvider, timer *mock.CircuitBreakerTimer, failures int) {
	failing := mp.Fail
	mp.Fail = true
	mp.TriedToSend = []mailspree.Message{}
	m := mailspree.Message{}
	for i := 0; i < 3; i++ {
		cb.SendEmail(m)
	}
	if len(mp.TriedToSend) != 3-failures {
		t.Errorf("the circuit breaker should be in closed state with %v failures", failures)
	}
	goToClosed(cb, mp, timer, failures)
	mp.Fail = failing
}

func goToClosed(cb mailspree.MailingProvider, mp *mock.MailingProvider, timer *mock.CircuitBreakerTimer, failures int) {
	// Ensure open state
	mp.Fail = true
	m := mailspree.Message{}
	for i := 0; i < 3; i++ {
		cb.SendEmail(m)
	}
	// Go to closed state with specific number of failures
	mp.Fail = false
	timer.EndTimer()
	time.Sleep(time.Millisecond * 200)
	cb.SendEmail(m)
	mp.Fail = true
	for i := 0; i < failures; i++ {
		cb.SendEmail(m)
	}
}

func testThatWeAreInOpen(t *testing.T, cb mailspree.MailingProvider, mp *mock.MailingProvider) {
	mp.TriedToSend = []mailspree.Message{}
	m := mailspree.Message{}
	cb.SendEmail(m)
	if len(mp.TriedToSend) != 0 {
		t.Error("the circuit breaker should be in opened state")
	}
}

func testThatWeAreInHalfOpen(t *testing.T, cb mailspree.MailingProvider, mp *mock.MailingProvider) {
	mp.Fail = true
	mp.TriedToSend = []mailspree.Message{}
	m := mailspree.Message{}
	cb.SendEmail(m)
	testThatWeAreInOpen(t, cb, mp)
}

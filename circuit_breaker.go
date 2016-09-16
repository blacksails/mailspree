package mailspree

import (
	"errors"
	"time"
)

// NewCircuitBreaker wraps a mailing provider with a circuit breaker.
func NewCircuitBreaker(mp MailingProvider, t CircuitBreakerTimer) MailingProvider {
	return &circuitBreaker{
		provider:            mp,
		state:               circuitBreakerClosed,
		failures:            0,
		failureLimit:        3,
		circuitBreakerTimer: t,
	}
}

// CircuitBreaker is a wrapper for mailing provides which provides circuit
// breaking functionality.
type circuitBreaker struct {
	provider            MailingProvider
	state               circuitBreakerState
	failures            int
	failureLimit        int
	circuitBreakerTimer CircuitBreakerTimer
}

// CircuitBreakerTimer is an interface to the timer used when the channel is open.
// This is defined as an interface so that we can make a mock implementation
type CircuitBreakerTimer interface {
	Run() <-chan int
}

type circuitBreakerTimer struct{}

func (t circuitBreakerTimer) Run() <-chan int {
	c := make(chan int)
	timer := time.NewTimer(time.Duration(30) * time.Second)
	go func() {
		<-timer.C
		c <- 0
	}()
	return c
}

// SendEmail tries to send the email with the mailing provider if the circuit
// is open.
func (cb *circuitBreaker) SendEmail(m Message) error {
	switch cb.state {
	case circuitBreakerClosed:
		err := cb.provider.SendEmail(m)
		if err != nil {
			cb.failures++
			if cb.failures >= cb.failureLimit {
				cb.switchState(circuitBreakerOpen)
			}
		} else {
			cb.failures = 0
		}
		return err
	case circuitBreakerOpen:
		return errors.New("Circuit breaker is open")
	case circuitBreakerHalfOpen:
		err := cb.provider.SendEmail(m)
		if err != nil {
			cb.switchState(circuitBreakerOpen)
		} else {
			cb.switchState(circuitBreakerClosed)
		}
		return err
	}
	return nil // we never get here
}

type circuitBreakerState int

const (
	circuitBreakerClosed circuitBreakerState = iota
	circuitBreakerOpen
	circuitBreakerHalfOpen
)

func (cb *circuitBreaker) switchState(s circuitBreakerState) {
	switch s {
	case circuitBreakerClosed:
		cb.failures = 0
	case circuitBreakerOpen:
		c := cb.circuitBreakerTimer.Run()
		go func() {
			<-c
			cb.switchState(circuitBreakerHalfOpen)
		}()
		cb.state = s
	}
	cb.state = s
}

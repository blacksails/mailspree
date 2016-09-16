package mailspree

import (
	"errors"
	"time"
)

// NewCircuitBreaker wraps a mailing provider with a circuit breaker.
func NewCircuitBreaker(mp MailingProvider, t CircuitBreakerTimer) MailingProvider {
	return circuitBreaker{
		provider:            mp,
		state:               circuitClosed,
		failures:            0,
		failureLimit:        3,
		circuitBreakerTimer: t,
	}
}

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
func (cb circuitBreaker) SendEmail(m Message) error {
	var err error
	switch cb.state {
	case circuitClosed:
		err = cb.provider.SendEmail(m)
		if err != nil {
			cb.failures++
			if cb.failures >= cb.failureLimit {
				cb.switchState(circuitOpen)
			}
		}
		return err
	case circuitOpen:
		return errors.New("Circuit breaker is open")
	case circuitHalfOpen:
		err = cb.provider.SendEmail(m)
		if err != nil {
			cb.switchState(circuitOpen)
		} else {
			cb.switchState(circuitClosed)
		}
		return err
	}
	return err
}

type circuitBreakerState int

const (
	circuitClosed circuitBreakerState = iota
	circuitOpen
	circuitHalfOpen
)

func (cb *circuitBreaker) switchState(s circuitBreakerState) {
	switch s {
	case circuitClosed:
		cb.failures = 0
	case circuitOpen:
		c := cb.circuitBreakerTimer.Run()
		go func() {
			<-c
			cb.switchState(circuitHalfOpen)
		}()
		cb.state = s
	}
	cb.state = s
}

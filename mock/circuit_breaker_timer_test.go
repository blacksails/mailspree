package mock_test

import (
	"testing"

	"github.com/blacksails/mailspree/mock"
)

func TestCircuitBreakerTimer(t *testing.T) {
	timer := mock.NewCircuitBreakerTimer()
	c := timer.Run()
	i := 42
	cont := make(chan int)
	go func() {
		i = <-c
		cont <- 0
	}()
	if i == 0 {
		t.Error("the timer should not have ended yet")
	}
	timer.EndTimer()
	<-cont
	if i != 0 {
		t.Error("the timer should have ended now")
	}
}

package mock

// CircuitBreakerTimer is a mock implementation of mailspree.CircuitBreakerTimer
type circuitBreakerTimer struct {
	c chan int
}

func NewCircuitBreakerTimer() *circuitBreakerTimer {
	return &circuitBreakerTimer{}
}

// Run returns a channel and it also sets it in the CircuitBreakerTimer struct,
// so that the timer can be ended using EndTimer.
func (t *circuitBreakerTimer) Run() <-chan int {
	c := make(chan int)
	t.c = c
	return c
}

// EndTimer sends a signal that the timer has ended.
func (t *circuitBreakerTimer) EndTimer() {
	t.c <- 0
}

package mock

// CircuitBreakerTimer is a mock implementation of mailspree.CircuitBreakerTimer
type CircuitBreakerTimer struct {
	c chan int
}

func NewCircuitBreakerTimer() *CircuitBreakerTimer {
	return &CircuitBreakerTimer{}
}

// Run returns a channel and it also sets it in the CircuitBreakerTimer struct,
// so that the timer can be ended using EndTimer.
func (t *CircuitBreakerTimer) Run() <-chan int {
	c := make(chan int)
	t.c = c
	return c
}

// EndTimer sends a signal that the timer has ended.
func (t *CircuitBreakerTimer) EndTimer() {
	t.c <- 0
}

package internal

// Root path Response
type FibonacciNumber struct {
	Number int
}

// sequence path JSON Response
type FibonacciNumberSequence struct {
	Numbers []int
}

// Error JSON Response
type EchoError struct {
	Message error
	Code    int
}

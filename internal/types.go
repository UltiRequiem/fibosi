package internal

// Root path Response
type FibonacciNumber struct {
	Number int `json:"number"`
}

// sequence path JSON Response
type FibonacciNumberSequence struct {
	Numbers []int `json:"numbers"`
}

// Error JSON Response
type EchoError struct {
	Message error `json:"message"`
	Code    int   `json:"code"`
}

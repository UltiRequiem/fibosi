package fibonacci

import "fmt"

var cache = make(map[int]int)

// Get Nth Fibonacci Number
func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("Got '%d' but expected  a number bigger than 0.", n)
	}

	if v, exits := cache[n]; exits {
		return v, nil
	}

	if n < 2 {
		cache[n] = n
		return n, nil
	}

	vOne, _ := Fibonacci(n - 2)

	vTwo, _ := Fibonacci(n - 1)

	v := vOne + vTwo

	cache[n] = v

	return v, nil
}

// Get an Slice with the first N numbers of the Fibonacci Sequence
func FibonacciSequence(length int) ([]int, error) {
	sequence := make([]int, length)

	for i := 0; i < length; i++ {
		fibo, err := Fibonacci(i)

		if err != nil {
			return make([]int, 0), err

		}

		sequence[i] = fibo
	}

	return sequence, nil
}

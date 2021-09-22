package fibonacci

import "fmt"

var cache = make(map[int]int)

func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("Got '%d' but expected  a number bigger than 0.", n)
	}

	// If already in cache, don't operate again
	v, exits := cache[n]

	if exits {
		return v, nil
	}

	// Equal
	if n < 2 {
		cache[n] = n
		return n, nil
	}

	// TODO: Use a loop instead of recursion here
	// I'dont check the error because I alredy know that is bigger than 0 :D
	vOne, _ := Fibonacci(n - 2)

	vTwo, _ := Fibonacci(n - 1)

	v = vOne + vTwo

	// Save to cache
	cache[n] = v

	return v, nil
}

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

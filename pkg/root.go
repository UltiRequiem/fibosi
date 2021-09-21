package fibonnaci

func Fibonacci(n int) int {
	f := make([]int, n+1, n+2)

	if n < 2 {
		f = f[0:2]
	}

	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}

func FibonacciSequence(n int) []int {
	var sequence []int

	for i := 0; i <= n; i++ {
		sequence = append(sequence, Fibonacci(i))
	}

	return sequence

}

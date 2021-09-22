package main

import (
	"fmt"

	fina "github.com/UltiRequiem/fibonacci/pkg"
)

func main() {
	fiboNum, _ := fina.Fibonacci(9)
	fmt.Println(fiboNum) // 34

	fiboSequence, _ := fina.FibonacciSequence(9)
	fmt.Println(fiboSequence) // [0 1 1 2 3 5 8 13 21]
}

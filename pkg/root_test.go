package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input, expectedOutput int
	}{
		{9, 34},
		{4, 3},
	}

	for _, test := range tests {
		result, _ := Fibonacci(test.input)

		if result != test.expectedOutput {
			t.Errorf("Incorrect result, got %d expected %d.", result, test.expectedOutput)
		}

	}
}

func TestFibonacciSequence(t *testing.T) {
	tests := []struct {
		input          int
		expectedOutput []int
	}{
		{8, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
	}

	for _, test := range tests {
		result, _ := FibonacciSequence(test.input)

		if reflect.DeepEqual(result, test.expectedOutput) {
			t.Errorf("Incorrect result, got %v expected %v.", result, test.expectedOutput)
		}

	}
}

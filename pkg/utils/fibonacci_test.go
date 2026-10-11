package utils

import (
	"testing"
)

func TestFibonacciValidCases(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{20, 6765},
		{30, 832040},
	}

	for _, testCase := range tests {
		actual, err := Fibonacci(testCase.input)
		if err != nil {
			t.Errorf("unexpected error for input %d: %v", testCase.input, err)
		}
		if actual != testCase.expected {
			t.Errorf("Fibonacci(%d) = %d; expected %d", testCase.input, actual, testCase.expected)
		}
	}
}

func TestFibonacciInvalidCases(t *testing.T) {
	invalidInputs := []int{-1, -10, 31, 100}

	for _, input := range invalidInputs {
		_, err := Fibonacci(input)
		if err == nil {
			t.Errorf("expected error for out-of-range input %d, got nil", input)
		}
	}
}

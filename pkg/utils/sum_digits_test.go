package utils

import (
	"testing"
)

func TestSumDigitsNormal(t *testing.T) {
	tests := []struct {
		input    int
		expected int
		hasError bool
	}{
		{input: 0, expected: 0, hasError: false},
		{input: 5, expected: 5, hasError: false},
		{input: 123, expected: 6, hasError: false},
		{input: 999, expected: 27, hasError: false},
		{input: 1045, expected: 10, hasError: false},
		{input: -1, expected: 0, hasError: true},
		{input: -123, expected: 0, hasError: true},
	}

	for _, testCase := range tests {
		actual, err := SumDigits(testCase.input)
		if testCase.hasError {
			if err == nil {
				t.Errorf("SumDigits(%d) expected error, got nil", testCase.input)
			}
		} else {
			if err != nil {
				t.Errorf("SumDigits(%d) unexpected error: %v", testCase.input, err)
			}
			if actual != testCase.expected {
				t.Errorf("SumDigits(%d) = %d, expected %d", testCase.input, actual, testCase.expected)
			}
		}
	}
}

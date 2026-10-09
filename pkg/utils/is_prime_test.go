package utils

import (
	"testing"
)

func TestIsPrimeTable(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Negative number",
			input:    -5,
			expected: false,
		},
		{
			name:     "Zero",
			input:    0,
			expected: false,
		},
		{
			name:     "One",
			input:    1,
			expected: false,
		},
		{
			name:     "Two",
			input:    2,
			expected: true,
		},
		{
			name:     "Three",
			input:    3,
			expected: true,
		},
		{
			name:     "Four",
			input:    4,
			expected: false,
		},
		{
			name:     "Prime eleven",
			input:    11,
			expected: true,
		},
		{
			name:     "Composite twenty-five",
			input:    25,
			expected: false,
		},
		{
			name:     "Prime ninety-seven",
			input:    97,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsPrime(tc.input)
			if actual != tc.expected {
				t.Errorf("IsPrime(%d) = %v; expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}

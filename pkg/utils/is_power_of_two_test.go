package utils

import (
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Negative integer",
			input:    -8,
			expected: false,
		},
		{
			name:     "Zero",
			input:    0,
			expected: false,
		},
		{
			name:     "One (2^0)",
			input:    1,
			expected: true,
		},
		{
			name:     "Two (2^1)",
			input:    2,
			expected: true,
		},
		{
			name:     "Three (not power of two)",
			input:    3,
			expected: false,
		},
		{
			name:     "Four (2^2)",
			input:    4,
			expected: true,
		},
		{
			name:     "Sixteen (2^4)",
			input:    16,
			expected: true,
		},
		{
			name:     "Eighteen (not power of two)",
			input:    18,
			expected: false,
		},
		{
			name:     "Large power of two (2^30)",
			input:    1073741824,
			expected: true,
		},
		{
			name:     "Large non-power of two",
			input:    1073741825,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsPowerOfTwo(tc.input)
			if actual != tc.expected {
				t.Errorf("IsPowerOfTwo(%d) = %v; expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}

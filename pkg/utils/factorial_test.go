package utils

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    int64
		expectedErr bool
	}{
		{
			name:        "zero",
			input:       0,
			expected:    1,
			expectedErr: false,
		},
		{
			name:        "one",
			input:       1,
			expected:    1,
			expectedErr: false,
		},
		{
			name:        "small positive",
			input:       5,
			expected:    120,
			expectedErr: false,
		},
		{
			name:        "upper boundary",
			input:       12,
			expected:    479001600,
			expectedErr: false,
		},
		{
			name:        "below lower boundary",
			input:       -1,
			expected:    0,
			expectedErr: true,
		},
		{
			name:        "above upper boundary",
			input:       13,
			expected:    0,
			expectedErr: true,
		},
		{
			name:        "far above upper boundary",
			input:       100,
			expected:    0,
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Factorial(tc.input)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("Factorial(%d) error = %v, expectedErr %v", tc.input, err, tc.expectedErr)
			}
			if !tc.expectedErr && actual != tc.expected {
				t.Errorf("Factorial(%d) = %d, expected %d", tc.input, actual, tc.expected)
			}
		})
	}
}

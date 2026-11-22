package utils

import (
	"testing"
)

func TestGcd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Both zeros",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "First zero",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "Second zero",
			a:        12,
			b:        0,
			expected: 12,
		},
		{
			name:     "Coprime numbers",
			a:        17,
			b:        13,
			expected: 1,
		},
		{
			name:     "Standard case",
			a:        54,
			b:        24,
			expected: 6,
		},
		{
			name:     "One divides other",
			a:        48,
			b:        12,
			expected: 12,
		},
		{
			name:     "Same numbers",
			a:        7,
			b:        7,
			expected: 7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := Gcd(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("Gcd(%d, %d) = %d; expected %d", tc.a, tc.b, actual, tc.expected)
			}
		})
	}
}

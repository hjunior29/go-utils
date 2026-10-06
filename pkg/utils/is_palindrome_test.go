package utils

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "simple even palindrome",
			input:    "abba",
			expected: true,
		},
		{
			name:     "simple odd palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "non-palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "case sensitive mismatch",
			input:    "Racecar",
			expected: false,
		},
		{
			name:     "unicode palindrome",
			input:    "level",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsPalindrome(tt.input)
			if actual != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, expected %v", tt.input, actual, tt.expected)
			}
		})
	}
}

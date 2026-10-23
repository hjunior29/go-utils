package utils

import (
	"testing"
)

func TestNormalizeWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Only whitespace",
			input:    "   \t\n ",
			expected: "",
		},
		{
			name:     "No whitespace changes needed",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "Leading and trailing whitespace",
			input:    "   hello world   ",
			expected: "hello world",
		},
		{
			name:     "Consecutive internal whitespace",
			input:    "hello     world",
			expected: "hello world",
		},
		{
			name:     "Mixed whitespace types",
			input:    "\t\n hello \t world \n",
			expected: "hello world",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := NormalizeWhitespace(tc.input)
			if actual != tc.expected {
				t.Errorf("NormalizeWhitespace(%q) = %q, expected %q", tc.input, actual, tc.expected)
			}
		})
	}
}

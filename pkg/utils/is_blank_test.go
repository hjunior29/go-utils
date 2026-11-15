package utils

import (
	"testing"
)

func TestIsBlank(t *testing.T) {
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
			name:     "spaces only",
			input:    "   ",
			expected: true,
		},
		{
			name:     "tabs and newlines",
			input:    "\t\n\r ",
			expected: true,
		},
		{
			name:     "unicode whitespace",
			input:    "\u00A0\u2000",
			expected: true,
		},
		{
			name:     "normal text",
			input:    "hello",
			expected: false,
		},
		{
			name:     "text with leading whitespace",
			input:    "  hello",
			expected: false,
		},
		{
			name:     "text with trailing whitespace",
			input:    "hello  ",
			expected: false,
		},
		{
			name:     "text with internal whitespace",
			input:    "hello world",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsBlank(tc.input)
			if actual != tc.expected {
				t.Errorf("IsBlank(%q) = %v; expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}

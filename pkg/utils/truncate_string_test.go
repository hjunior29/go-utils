package utils

import (
	"testing"
)

func TestTruncateString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxRunes int
		expected string
	}{
		{
			name:     "normal truncation ASCII",
			input:    "hello world",
			maxRunes: 5,
			expected: "hello",
		},
		{
			name:     "normal truncation Unicode",
			input:    "こんにちは",
			maxRunes: 3,
			expected: "こんに",
		},
		{
			name:     "maxRunes greater than length",
			input:    "abc",
			maxRunes: 5,
			expected: "abc",
		},
		{
			name:     "maxRunes equal to length",
			input:    "abc",
			maxRunes: 3,
			expected: "abc",
		},
		{
			name:     "zero maxRunes",
			input:    "abc",
			maxRunes: 0,
			expected: "",
		},
		{
			name:     "empty string input",
			input:    "",
			maxRunes: 5,
			expected: "",
		},
		{
			name:     "negative maxRunes returns original",
			input:    "abc",
			maxRunes: -1,
			expected: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := TruncateString(tt.input, tt.maxRunes)
			if actual != tt.expected {
				t.Errorf("TruncateString(%q, %d) = %q; expected %q", tt.input, tt.maxRunes, actual, tt.expected)
			}
		})
	}
}

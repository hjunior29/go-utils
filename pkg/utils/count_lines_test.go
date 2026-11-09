package utils

import (
	"testing"
)

func TestCountLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Single line without newline",
			input:    "hello",
			expected: 1,
		},
		{
			name:     "Single line with trailing newline",
			input:    "hello\n",
			expected: 1,
		},
		{
			name:     "Multiple lines without trailing newline",
			input:    "line1\nline2\nline3",
			expected: 3,
		},
		{
			name:     "Multiple lines with trailing newline",
			input:    "line1\nline2\nline3\n",
			expected: 3,
		},
		{
			name:     "Consecutive newlines",
			input:    "\n\n\n",
			expected: 3,
		},
		{
			name:     "Single newline",
			input:    "\n",
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := CountLines(tc.input)
			if actual != tc.expected {
				t.Errorf("CountLines(%q) = %d; expected %d", tc.input, actual, tc.expected)
			}
		})
	}
}

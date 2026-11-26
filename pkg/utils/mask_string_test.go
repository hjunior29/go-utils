package utils

import (
	"testing"
)

func TestMaskString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		n           int
		expected    string
		expectError bool
	}{
		{
			name:        "Normal masking",
			input:       "1234567890",
			n:           4,
			expected:    "******7890",
			expectError: false,
		},
		{
			name:        "Zero N",
			input:       "secret",
			n:           0,
			expected:    "******",
			expectError: false,
		},
		{
			name:        "N greater than length",
			input:       "abc",
			n:           5,
			expected:    "abc",
			expectError: false,
		},
		{
			name:        "N equal to length",
			input:       "abc",
			n:           3,
			expected:    "abc",
			expectError: false,
		},
		{
			name:        "Empty string",
			input:       "",
			n:           2,
			expected:    "",
			expectError: false,
		},
		{
			name:        "Unicode code points",
			input:       "héllo",
			n:           2,
			expected:    "***lo",
			expectError: false,
		},
		{
			name:        "Negative N error",
			input:       "password",
			n:           -1,
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := MaskString(tc.input, tc.n)
			if tc.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

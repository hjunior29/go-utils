package utils

import (
	"testing"
)

func TestCommonSuffixNormal(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{"programming", "coding", "ing"},
		{"hello", "world", ""},
		{"same", "same", "same"},
		{"", "abc", ""},
		{"abc", "", ""},
		{"", "", ""},
		{"日本語", "本語", "本語"},
	}

	for _, tt := range tests {
		actual := CommonSuffix(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("CommonSuffix(%q, %q) = %q, expected %q", tt.a, tt.b, actual, tt.expected)
		}
	}
}

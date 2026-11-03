package utils

import (
	"testing"
)

func TestCommonPrefixNormal(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{"interspecies", "interstellar", "inters"},
		{"apple", "apricot", "ap"},
		{"abcdef", "abcdef", "abcdef"},
		{"abc", "xyz", ""},
		{"", "abc", ""},
		{"abc", "", ""},
		{"", "", ""},
		{"こんにちは世界", "こんにちは日本", "こんにちは"},
	}

	for _, tc := range tests {
		actual := CommonPrefix(tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("CommonPrefix(%q, %q) = %q; expected %q", tc.a, tc.b, actual, tc.expected)
		}
	}
}

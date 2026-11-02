package utils

import (
	"testing"
)

func TestHammingDistanceValid(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{
			name:     "empty strings",
			a:        "",
			b:        "",
			expected: 0,
		},
		{
			name:     "identical ASCII strings",
			a:        "karolin",
			b:        "karolin",
			expected: 0,
		},
		{
			name:     "different ASCII strings",
			a:        "karolin",
			b:        "kathrin",
			expected: 3,
		},
		{
			name:     "identical Unicode strings",
			a:        "café",
			b:        "café",
			expected: 0,
		},
		{
			name:     "different Unicode strings",
			a:        "café",
			b:        "cage",
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := HammingDistance(tc.a, tc.b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestHammingDistanceInvalid(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
	}{
		{
			name: "different lengths ASCII",
			a:    "abc",
			b:    "abcd",
		},
		{
			name: "empty vs non-empty",
			a:    "",
			b:    "a",
		},
		{
			name: "different rune counts same byte lengths",
			a:    "a",
			b:    "ab",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := HammingDistance(tc.a, tc.b)
			if err == nil {
				t.Error("expected error for unequal lengths, got nil")
			}
		})
	}
}

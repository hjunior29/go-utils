package utils

import (
	"reflect"
	"testing"
)

func TestDeduplicateInts_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "normal case with duplicates",
			input:    []int{1, 2, 2, 3, 1, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name:     "no duplicates",
			input:    []int{5, 6, 7},
			expected: []int{5, 6, 7},
		},
		{
			name:     "all duplicates",
			input:    []int{9, 9, 9, 9},
			expected: []int{9},
		},
		{
			name:     "negative numbers and zero",
			input:    []int{-1, 0, -1, 2, 0},
			expected: []int{-1, 0, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := DeduplicateInts(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("DeduplicateInts(%v) = %v; expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}

package utils

import (
	"reflect"
	"testing"
)

func TestIntersectionInts(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "Standard intersection with duplicates",
			first:    []int{1, 2, 3, 2, 4},
			second:   []int{2, 3, 5},
			expected: []int{2, 3},
		},
		{
			name:     "Preserve first sequence order",
			first:    []int{5, 1, 3, 1, 2},
			second:   []int{1, 2, 3, 5},
			expected: []int{5, 1, 3, 2},
		},
		{
			name:     "No common elements",
			first:    []int{1, 2, 3},
			second:   []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "First slice empty",
			first:    []int{},
			second:   []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Second slice empty",
			first:    []int{1, 2, 3},
			second:   []int{},
			expected: []int{},
		},
		{
			name:     "Both slices empty",
			first:    []int{},
			second:   []int{},
			expected: []int{},
		},
		{
			name:     "Nil slices",
			first:    nil,
			second:   nil,
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := IntersectionInts(tc.first, tc.second)
			if len(actual) == 0 && len(tc.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("IntersectionInts(%v, %v) = %v, expected %v", tc.first, tc.second, actual, tc.expected)
			}
		})
	}
}

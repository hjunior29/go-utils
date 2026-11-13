package utils

import (
	"reflect"
	"testing"
)

func TestChunkInts(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		size        int
		expected    [][]int
		expectError bool
	}{
		{
			name:        "normal division into equal chunks",
			input:       []int{1, 2, 3, 4, 5, 6},
			size:        2,
			expected:    [][]int{{1, 2}, {3, 4}, {5, 6}},
			expectError: false,
		},
		{
			name:        "uneven chunks with remainder",
			input:       []int{1, 2, 3, 4, 5},
			size:        2,
			expected:    [][]int{{1, 2}, {3, 4}, {5}},
			expectError: false,
		},
		{
			name:        "chunk size larger than input",
			input:       []int{1, 2, 3},
			size:        5,
			expected:    [][]int{{1, 2, 3}},
			expectError: false,
		},
		{
			name:        "exact single chunk",
			input:       []int{1, 2, 3},
			size:        3,
			expected:    [][]int{{1, 2, 3}},
			expectError: false,
		},
		{
			name:        "empty input slice",
			input:       []int{},
			size:        3,
			expected:    [][]int{},
			expectError: false,
		},
		{
			name:        "nil input slice",
			input:       nil,
			size:        2,
			expected:    [][]int{},
			expectError: false,
		},
		{
			name:        "zero chunk size rejects",
			input:       []int{1, 2, 3},
			size:        0,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "negative chunk size rejects",
			input:       []int{1, 2, 3},
			size:        -1,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ChunkInts(tc.input, tc.size)

			if tc.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("ChunkInts(%v, %d) = %v, expected %v", tc.input, tc.size, actual, tc.expected)
			}
		})
	}
}

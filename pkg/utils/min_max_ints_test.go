package utils

import (
	"testing"
)

func TestMinMaxInts(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expectedMin int
		expectedMax int
		expectError bool
	}{
		{
			name:        "empty slice",
			input:       []int{},
			expectedMin: 0,
			expectedMax: 0,
			expectError: true,
		},
		{
			name:        "single element",
			input:       []int{42},
			expectedMin: 42,
			expectedMax: 42,
			expectError: false,
		},
		{
			name:        "multiple positive elements",
			input:       []int{5, 2, 9, 1, 7},
			expectedMin: 1,
			expectedMax: 9,
			expectError: false,
		},
		{
			name:        "mixed positive and negative elements",
			input:       []int{-3, 10, 0, -5, 8},
			expectedMin: -5,
			expectedMax: 10,
			expectError: false,
		},
		{
			name:        "all negative elements",
			input:       []int{-10, -2, -50, -5},
			expectedMin: -50,
			expectedMax: -2,
			expectError: false,
		},
		{
			name:        "duplicate min and max elements",
			input:       []int{3, 3, 1, 1, 5, 5},
			expectedMin: 1,
			expectedMax: 5,
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			minVal, maxVal, err := MinMaxInts(tc.input)
			if tc.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if minVal != tc.expectedMin {
					t.Errorf("expected min %d, got %d", tc.expectedMin, minVal)
				}
				if maxVal != tc.expectedMax {
					t.Errorf("expected max %d, got %d", tc.expectedMax, maxVal)
				}
			}
		})
	}
}

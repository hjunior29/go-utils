package utils

import (
	"reflect"
	"testing"
)

func TestPartitionInts_Standard(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	wantedEvens := []int{2, 4, 6}
	wantedOdds := []int{1, 3, 5}

	evens, odds := PartitionInts(input)

	if !reflect.DeepEqual(evens, wantedEvens) {
		t.Errorf("PartitionInts() evens = %v, want %v", evens, wantedEvens)
	}
	if !reflect.DeepEqual(odds, wantedOdds) {
		t.Errorf("PartitionInts() odds = %v, want %v", odds, wantedOdds)
	}
}

func TestPartitionInts_EmptyAndNil(t *testing.T) {
	evens, odds := PartitionInts([]int{})
	if evens == nil || len(evens) != 0 {
		t.Errorf("PartitionInts() empty input should return non-nil empty slice for evens, got %v", evens)
	}
	if odds == nil || len(odds) != 0 {
		t.Errorf("PartitionInts() empty input should return non-nil empty slice for odds, got %v", odds)
	}

	evensNil, oddsNil := PartitionInts(nil)
	if evensNil != nil {
		t.Errorf("PartitionInts() nil input should return nil for evens, got %v", evensNil)
	}
	if oddsNil != nil {
		t.Errorf("PartitionInts() nil input should return nil for odds, got %v", oddsNil)
	}
}

func TestPartitionInts_TableDriven(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		wantedEvens []int
		wantedOdds  []int
	}{
		{
			name:        "All evens",
			input:       []int{2, 4, 6, 8},
			wantedEvens: []int{2, 4, 6, 8},
			wantedOdds:  []int{},
		},
		{
			name:        "All odds",
			input:       []int{1, 3, 5, 7},
			wantedEvens: []int{},
			wantedOdds:  []int{1, 3, 5, 7},
		},
		{
			name:        "Negative numbers",
			input:       []int{-3, -2, -1, 0, 1, 2, 3},
			wantedEvens: []int{-2, 0, 2},
			wantedOdds:  []int{-3, -1, 1, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			evens, odds := PartitionInts(tc.input)
			if !reflect.DeepEqual(evens, tc.wantedEvens) {
				t.Errorf("evens = %v, want %v", evens, tc.wantedEvens)
			}
			if !reflect.DeepEqual(odds, tc.wantedOdds) {
				t.Errorf("odds = %v, want %v", odds, tc.wantedOdds)
			}
		})
	}
}

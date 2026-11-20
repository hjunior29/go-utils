package utils

import (
	"reflect"
	"testing"
)

func TestUnionInts(t *testing.T) {
	tests := []struct {
		name   string
		first  []int
		second []int
		want   []int
	}{
		{
			name:   "standard disjoint",
			first:  []int{1, 2, 3},
			second: []int{4, 5, 6},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "overlapping elements",
			first:  []int{1, 2, 3, 2},
			second: []int{2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "both empty",
			first:  []int{},
			second: []int{},
			want:   nil,
		},
		{
			name:   "first empty",
			first:  []int{},
			second: []int{4, 5, 6},
			want:   []int{4, 5, 6},
		},
		{
			name:   "second empty",
			first:  []int{1, 2, 3},
			second: []int{},
			want:   []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionInts(tt.first, tt.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

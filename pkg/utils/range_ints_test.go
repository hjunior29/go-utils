package utils

import (
	"reflect"
	"testing"
)

func TestRangeInts_Table(t *testing.T) {
	tests := []struct {
		name       string
		start      int
		end        int
		step       int
		wantResult []int
		wantError  bool
	}{
		{
			name:       "positive step ascending",
			start:      0,
			end:        5,
			step:       1,
			wantResult: []int{0, 1, 2, 3, 4},
			wantError:  false,
		},
		{
			name:       "positive step ascending with custom step",
			start:      1,
			end:        10,
			step:       3,
			wantResult: []int{1, 4, 7},
			wantError:  false,
		},
		{
			name:       "positive step descending",
			start:      5,
			end:        0,
			step:       1,
			wantResult: []int{5, 4, 3, 2, 1},
			wantError:  false,
		},
		{
			name:       "empty range equal bounds",
			start:      3,
			end:        3,
			step:       1,
			wantResult: nil,
			wantError:  false,
		},
		{
			name:       "zero step invalid",
			start:      0,
			end:        5,
			step:       0,
			wantResult: nil,
			wantError:  true,
		},
		{
			name:       "negative step invalid",
			start:      0,
			end:        5,
			step:       -1,
			wantResult: nil,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RangeInts(tt.start, tt.end, tt.step)
			if (err != nil) != tt.wantError {
				t.Errorf("RangeInts() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResult) {
				t.Errorf("RangeInts() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

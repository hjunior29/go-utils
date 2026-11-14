package utils

import (
	"errors"
)

// MinMaxInts returns the minimum and maximum values of a nonempty integer sequence.
// It returns an error if the input slice is empty.
func MinMaxInts(values []int) (int, int, error) {
	if len(values) == 0 {
		return 0, 0, errors.New("empty integer sequence")
	}

	minVal := values[0]
	maxVal := values[0]

	for _, v := range values[1:] {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	return minVal, maxVal, nil
}

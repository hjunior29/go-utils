package utils

import (
	"errors"
)

// RangeInts produces integers from an inclusive start to exclusive end using a positive step.
// It returns an error if the step is less than or equal to zero.
func RangeInts(start, end, step int) ([]int, error) {
	if step <= 0 {
		return nil, errors.New("step must be positive")
	}

	var result []int
	if start < end {
		for current := start; current < end; current += step {
			result = append(result, current)
		}
	} else {
		for current := start; current > end; current -= step {
			result = append(result, current)
		}
	}

	return result, nil
}

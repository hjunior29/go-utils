package utils

import (
	"errors"
)

// TakeInts returns the first n integers from the input slice without changing the input.
// It returns an error if n is negative.
func TakeInts(values []int, n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("negative count is not allowed")
	}
	if n == 0 {
		return []int{}, nil
	}
	if n > len(values) {
		n = len(values)
	}

	result := make([]int, n)
	copy(result, values[:n])
	return result, nil
}

package utils

import (
	"errors"
)

// RotateInts rotates an integer sequence left by a nonnegative offset without changing the input.
func RotateInts(nums []int, offset int) ([]int, error) {
	if offset < 0 {
		return nil, errors.New("offset cannot be negative")
	}

	if len(nums) == 0 {
		result := make([]int, 0)
		return result, nil
	}

	effectiveOffset := offset % len(nums)
	if effectiveOffset == 0 {
		result := make([]int, len(nums))
		copy(result, nums)
		return result, nil
	}

	result := make([]int, len(nums))
	copy(result, nums[effectiveOffset:])
	copy(result[len(nums)-effectiveOffset:], nums[:effectiveOffset])
	return result, nil
}

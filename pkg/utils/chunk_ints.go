package utils

import (
	"errors"
)

// ChunkInts splits an integer slice into nonempty chunks of the given positive size.
// It returns an error if the size is less than or equal to zero.
func ChunkInts(input []int, size int) ([][]int, error) {
	if size <= 0 {
		return nil, errors.New("chunk size must be positive")
	}

	if len(input) == 0 {
		return [][]int{}, nil
	}

	var result [][]int
	for i := 0; i < len(input); i += size {
		end := i + size
		if end > len(input) {
			end = len(input)
		}

		chunk := make([]int, end-i)
		copy(chunk, input[i:end])
		result = append(result, chunk)
	}

	return result,
		nil
}

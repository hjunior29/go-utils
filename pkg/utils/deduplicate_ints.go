package utils

// DeduplicateInts removes duplicate integers while preserving first occurrence order.
func DeduplicateInts(input []int) []int {
	if input == nil {
		return nil
	}
	seen := make(map[int]bool)
	result := make([]int, 0, len(input))
	for _, val := range input {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}
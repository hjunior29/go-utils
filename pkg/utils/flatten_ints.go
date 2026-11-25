package utils

// FlattenInts flattens a sequence of integer sequences while preserving order.
func FlattenInts(sequences [][]int) []int {
	totalLen := 0
	for _, seq := range sequences {
		totalLen += len(seq)
	}

	result := make([]int, 0, totalLen)
	for _, seq := range sequences {
		result = append(result, seq...)
	}

	return result
}

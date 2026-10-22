package utils

// IntersectionInts returns unique shared integers in first-sequence order.
func IntersectionInts(first []int, second []int) []int {
	if len(first) == 0 || len(second) == 0 {
		return []int{}
	}

	lookup := make(map[int]bool, len(second))
	for _, val := range second {
		lookup[val] = true
	}

	seen := make(map[int]bool)
	var result []int

	for _, val := range first {
		if lookup[val] && !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}

	if result == nil {
		return []int{}
	}

	return result
}

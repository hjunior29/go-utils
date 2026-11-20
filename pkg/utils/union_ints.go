package utils

func UnionInts(first []int, second []int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, val := range first {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	for _, val := range second {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

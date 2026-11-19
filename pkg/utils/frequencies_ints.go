package utils

func FrequenciesInts(numbers []int) map[int]int {
	frequencies := make(map[int]int)
	for _, number := range numbers {
		frequencies[number]++
	}
	return frequencies
}

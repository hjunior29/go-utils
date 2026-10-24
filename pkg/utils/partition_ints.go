package utils

func PartitionInts(numbers []int) ([]int, []int) {
	if numbers == nil {
		return nil, nil
	}

	var evens []int
	var odds []int

	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		} else {
			odds = append(odds, n)
		}
	}

	if evens == nil {
		evens = []int{}
	}
	if odds == nil {
		odds = []int{}
	}

	return evens, odds
}

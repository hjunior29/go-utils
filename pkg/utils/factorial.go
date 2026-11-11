package utils

import (
	"fmt"
)

// Factorial computes the factorial for integers 0 through 12.
// It returns an error if the input is outside the range [0, 12].
func Factorial(n int) (int64, error) {
	if n < 0 || n > 12 {
		return 0, fmt.Errorf("input %d is out of range [0, 12]", n)
	}

	var result int64 = 1
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}

	return result, nil
}

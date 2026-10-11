package utils

import (
	"errors"
)

// Fibonacci returns the Fibonacci number F(n) for 0 through 30.
// It returns an error if n is outside the inclusive range [0, 30].
func Fibonacci(n int) (int, error) {
	if n < 0 || n > 30 {
		return 0, errors.New("input out of allowed range [0, 30]")
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

package utils

import (
	"errors"
)

func SumDigits(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("negative input is not supported")
	}
	if n == 0 {
		return 0, nil
	}
	sum := 0
	temp := n
	for temp > 0 {
		sum += temp % 10
		temp /= 10
	}
	return sum, nil
}

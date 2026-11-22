package utils

// Gcd computes the greatest common divisor of two nonnegative integers.
// If both inputs are zero, it returns zero.
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

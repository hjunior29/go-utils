package utils

// IsPowerOfTwo checks whether a positive integer is a power of two.
func IsPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

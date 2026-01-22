```go
package utils

import (
	"strings"
	"unicode"
)

// Reverse returns the reverse of a string.
// It handles Unicode characters correctly by operating on runes.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Capitalize returns the string with the first letter capitalized.
// If the string is empty, it returns an empty string.
// It handles Unicode characters correctly.
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Contains checks if a slice of strings contains a specific string.
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// TrimAll removes all whitespace characters (spaces, tabs, newlines, etc.) from a string.
// It uses a strings.Builder for efficient string concatenation.
func TrimAll(s string) string {
	var builder strings.Builder
	builder.Grow(len(s)) // Pre-allocate capacity for efficiency
	for _, r := range s {
		if !unicode.IsSpace(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// Max returns the maximum of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Filter returns a new slice containing only elements from the input slice
// that satisfy the given predicate function.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// Pre-allocate result slice if possible, though exact size is unknown.
	// A small initial capacity can still be beneficial.
	result := make([]T, 0, len(slice)/2) // Heuristic initial capacity
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Clamp restricts an integer value to be within a specified range [min, max].
// If val < min, it returns min. If val > max, it returns max.
// It assumes min <= max.
func Clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// Abs returns the absolute value of an integer.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IsEmpty checks if a string is empty or contains only whitespace.
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Truncate returns the first n characters of a string.
// If the string is shorter than n, the original string is returned.
// If n is negative, the original string is returned.
// This function operates on runes to correctly handle multi-byte characters.
func Truncate(s string, n int) string {
	if n < 0 {
		return s
	}
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n])
}

// IsPalindrome checks if a string is a palindrome (reads the same forwards and backwards).
// It handles Unicode characters correctly. Case-insensitive.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s) // Case-insensitive comparison
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// Repeat returns a new string consisting of n copies of the string s.
// If n is zero or negative, an empty string is returned.
func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}
```
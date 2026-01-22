package utils

import (
	"errors"
	"strings"
	"unicode"
)

// Reverse returns the reverse of a string.
// It handles Unicode characters correctly by operating on runes.
// For example:
//   Reverse("hello") == "olleh"
//   Reverse("你好") == "好你"
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
// For example:
//   Capitalize("hello") == "Hello"
//   Capitalize("你好") == "你好" (if the first character has no uppercase form)
//   Capitalize("") == ""
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Contains checks if a slice of strings contains a specific string.
// For example:
//   Contains([]string{"a", "b", "c"}, "b") == true
//   Contains([]string{"a", "b", "c"}, "d") == false
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
// For example:
//   TrimAll("  hello \t world \n") == "helloworld"
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
// For example:
//   Max(5, 10) == 10
//   Max(10, 5) == 10
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two integers.
// For example:
//   Min(5, 10) == 5
//   Min(10, 5) == 5
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Filter returns a new slice containing only elements from the input slice
// that satisfy the given predicate function.
// The predicate function should return true for elements to keep and false for elements to discard.
// For example:
//   Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }) == []int{2, 4}
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// Pre-allocate result slice with a heuristic initial capacity.
	// This can improve performance by reducing reallocations.
	result := make([]T, 0, len(slice)/2)
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Clamp restricts an integer value to be within a specified range [min, max].
// If val < min, it returns min. If val > max, it returns max.
// It returns an error if min > max.
// For example:
//   Clamp(5, 0, 10) == 5
//   Clamp(-5, 0, 10) == 0
//   Clamp(15, 0, 10) == 10
//   Clamp(5, 10, 0) returns an error
func Clamp(val, min, max int) (int, error) {
	if min > max {
		return 0, errors.New("min cannot be greater than max")
	}
	if val < min {
		return min, nil
	}
	if val > max {
		return max, nil
	}
	return val, nil
}

// Abs returns the absolute value of an integer.
// For example:
//   Abs(5) == 5
//   Abs(-5) == 5
//   Abs(0) == 0
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IsEmpty checks if a string is empty or contains only whitespace.
// It uses strings.TrimSpace to remove leading/trailing whitespace before checking.
// For example:
//   IsEmpty("") == true
//   IsEmpty("   ") == true
//   IsEmpty("hello") == false
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Truncate returns the first n characters of a string.
// If the string is shorter than n, the original string is returned.
// If n is negative, the original string is returned.
// This function operates on runes to correctly handle multi-byte characters.
// For example:
//   Truncate("hello world", 5) == "hello"
//   Truncate("你好世界", 2) == "你好"
//   Truncate("short", 10) == "short"
//   Truncate("string", -1) == "string"
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
// It handles Unicode characters correctly and performs a case-insensitive comparison.
// It ignores non-alphanumeric characters.
// For example:
//   IsPalindrome("madam") == true
//   IsPalindrome("Madam") == true
//   IsPalindrome("hello") == false
//   IsPalindrome("racecar") == true
//   IsPalindrome("A man, a plan, a canal: Panama") == true
func IsPalindrome(s string) bool {
	var cleanedRunes []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleanedRunes = append(cleanedRunes, unicode.ToLower(r))
		}
	}

	for i, j := 0, len(cleanedRunes)-1; i < j; i, j = i+1, j-1 {
		if cleanedRunes[i] != cleanedRunes[j] {
			return false
		}
	}
	return true
}

// Repeat returns a new string consisting of n copies of the string s.
// If n is zero or negative, an empty string is returned.
// For example:
//   Repeat("abc", 3) == "abcabcabc"
//   Repeat("a", 5) == "aaaaa"
//   Repeat("abc", 0) == ""
//   Repeat("abc", -2) == ""
func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	// strings.Repeat is efficient, but for very large n, a builder might offer
	// a slight advantage in some scenarios by pre-allocating.
	// However, strings.Repeat is generally well-optimized.
	return strings.Repeat(s, n)
}

// ContainsAny checks if a string contains any of the characters from a given set of runes.
// For example:
//   ContainsAny("hello", []rune{'a', 'e', 'i'}) == true
//   ContainsAny("world", []rune{'a', 'e', 'i'}) == false
func ContainsAny(s string, runes []rune) bool {
	runeSet := make(map[rune]struct{}, len(runes))
	for _, r := range runes {
		runeSet[r]
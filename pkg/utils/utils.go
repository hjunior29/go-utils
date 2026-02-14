package utils

import (
	"errors"
	"strings"
	"unicode"
)

// Reverse returns the reverse of a string.
// It handles Unicode characters correctly by operating on runes.
//
// Examples:
//
//	Reverse("hello") == "olleh"
//	Reverse("你好") == "好你"
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
//
// Examples:
//
//	Capitalize("hello") == "Hello"
//	Capitalize("你好") == "你好" (if the first character has no uppercase form)
//	Capitalize("") == ""
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Contains checks if a slice of strings contains a specific string.
//
// Examples:
//
//	Contains([]string{"a", "b", "c"}, "b") == true
//	Contains([]string{"a", "b", "c"}, "d") == false
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
//
// Examples:
//
//	TrimAll("  hello \t world \n") == "helloworld"
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
//
// Examples:
//
//	Max(5, 10) == 10
//	Max(10, 5) == 10
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two integers.
//
// Examples:
//
//	Min(5, 10) == 5
//	Min(10, 5) == 5
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Filter returns a new slice containing only elements from the input slice
// that satisfy the given predicate function.
// The predicate function should return true for elements to keep and false for elements to discard.
//
// Examples:
//
//	Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }) == []int{2, 4}
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
//
// Examples:
//
//	Clamp(5, 0, 10) == 5
//	Clamp(-5, 0, 10) == 0
//	Clamp(15, 0, 10) == 10
//	Clamp(5, 10, 0) returns an error
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
//
// Examples:
//
//	Abs(5) == 5
//	Abs(-5) == 5
//	Abs(0) == 0
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IsEmpty checks if a string is empty or contains only whitespace.
// It uses strings.TrimSpace to remove leading/trailing whitespace before checking.
//
// Examples:
//
//	IsEmpty("") == true
//	IsEmpty("   ") == true
//	IsEmpty("hello") == false
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Truncate returns the first n characters of a string.
// If the string is shorter than n, the original string is returned.
// If n is negative, the original string is returned.
// This function operates on runes to correctly handle multi-byte characters.
//
// Examples:
//
//	Truncate("hello world", 5) == "hello"
//	Truncate("你好世界", 2) == "你好"
//	Truncate("short", 10) == "short"
//	Truncate("string", -1) == "string"
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
//
// Examples:
//
//	IsPalindrome("madam") == true
//	IsPalindrome("Madam") == true
//	IsPalindrome("hello") == false
//	IsPalindrome("racecar") == true
//	IsPalindrome("A man, a plan, a canal: Panama") == true
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
//
// Examples:
//
//	Repeat("abc", 3) == "abcabcabc"
//	Repeat("a", 5) == "aaaaa"
//	Repeat("abc", 0) == ""
//	Repeat("abc", -2) == ""
func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}

// ContainsAny checks if a string contains any of the characters from a given set of runes.
//
// Examples:
//
//	ContainsAny("hello", []rune{'a', 'e', 'i'}) == true
//	ContainsAny("world", []rune{'a', 'e', 'i'}) == false
func ContainsAny(s string, chars []rune) bool {
	if len(s) == 0 || len(chars) == 0 {
		return false
	}
	// Create a map for efficient lookup of characters to check.
	charSet := make(map[rune]struct{}, len(chars))
	for _, char := range chars {
		charSet[char] = struct{}{}
	}

	for _, r := range s {
		if _, ok := charSet[r]; ok {
			return true
		}
	}
	return false
}

// Slugify converts a string into a URL-friendly slug.
// It converts the string to lowercase, replaces spaces and non-alphanumeric characters with hyphens,
// and trims leading/trailing hyphens. Multiple hyphens are reduced to a single hyphen.
//
// Examples:
//
//	Slugify("Hello World!") == "hello-world"
//	Slugify(" A New Topic  ") == "a-new-topic"
//	Slugify("Another_Example-Here") == "another-example-here"
//	Slugify("123-456") == "123-456"
func Slugify(s string) string {
	s = strings.ToLower(s)
	var builder strings.Builder
	var lastCharIsHyphen bool

	for i, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder.WriteRune(r)
			lastCharIsHyphen = false
		} else if !lastCharIsHyphen {
			// Only add a hyphen if it's not a duplicate and not at the beginning
			if builder.Len() > 0 && i < len(s) {
				builder.WriteRune('-')
				lastCharIsHyphen = true
			}
		}
	}

	// Trim leading and trailing hyphens
	result := builder.String()
	result = strings.Trim(result, "-")
	return result
}
// ValidateLength checks if a string's length is within a specified range.
// It returns an error if the length is less than minLength or greater than maxLength.
// If minLength is negative or maxLength is negative, it's considered invalid.
// If minLength > maxLength, it's considered invalid.
//
// Examples:
//
//	ValidateLength("hello", 3, 5) == nil
//	ValidateLength("hi", 3, 5) returns an error
//	ValidateLength("hello world", 3, 5) returns an error
//	ValidateLength("test", -1, 5) returns an error
//	ValidateLength("test", 5, -1) returns an error
//	ValidateLength("test", 5, 3) returns an error
func ValidateLength(s string, minLength, maxLength int) error {
	if minLength < 0 || maxLength < 0 {
		return errors.New("minLength and maxLength must be non-negative")
	}
	if minLength > maxLength {
		return errors.New("minLength cannot be greater than maxLength")
	}
	length := len([]rune(s)) // Use runes to handle multi-byte characters
	if length < minLength {
		return errors.New("string is too short")
	}
	if length > maxLength {
		return errors.New("string is too long")
	}
	return nil
}

// ToTitleCase converts a string to title case, capitalizing the first letter of each word.
// Words are delimited by spaces. It handles Unicode characters correctly.
//
// Examples:
//
//	ToTitleCase("hello world") == "Hello World"
//	ToTitleCase("a song of ice and fire") == "A Song Of Ice And Fire"
//	ToTitleCase("HELLO WORLD") == "Hello World"
//	ToTitleCase("") == ""
//	ToTitleCase("  leading spaces") == "  Leading Spaces"
func ToTitleCase(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	capitalizeNext := true
	for i, r := range runes {
		if unicode.IsSpace(r) {
			capitalizeNext = true
		} else if capitalizeNext {
			runes[i] = unicode.ToUpper(r)
			capitalizeNext = false
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}

// SafeClamp restricts an integer value to be within a specified range [min, max].
// If val < min, it returns min. If val > max, it returns max.
// It returns an error if min > max.
//
// Examples:
//
//	SafeClamp(5, 0, 10) == (5, nil)
//	SafeClamp(-5, 0, 10) == (0, nil)
//	SafeClamp(15, 0, 10) == (10, nil)
//	SafeClamp(5, 10, 0) returns (0, error)
func SafeClamp(val, min, max int) (int, error) {
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

// SafeTruncate returns the first n characters of a string.
// If the string is shorter than n, the original string is returned.
// If n is negative, an error is returned.
// This function operates on runes to correctly handle multi-byte characters.
//
// Examples:
//
//	SafeTruncate("hello world", 5) == ("hello", nil)
//	SafeTruncate("你好世界", 2) == ("你好", nil)
//	SafeTruncate("short", 10) == ("short", nil)
//	SafeTruncate("string", -1) returns ("", error)
func SafeTruncate(s string, n int) (string, error) {
	if n < 0 {
		return "", errors.New("n cannot be negative")
	}
	runes := []rune(s)
	if len(runes) <= n {
		return s, nil
	}
	return string(runes[:n]), nil
}

// SafeIndex returns the index of the first instance of substr in s, or -1 if substr is not present in s.
// If substr is empty, it returns 0.
// It returns an error if the substring is not found.
//
// Examples:
//
//	SafeIndex("hello world", "world") == (6, nil)
//	SafeIndex("hello world", "goodbye") returns (-1, error)
//	SafeIndex("hello", "") == (0, nil)
func SafeIndex(s, substr string) (int, error) {
	index := strings.Index(s, substr)
	if index == -1 {
		return -1, errors.New("substring not found")
	}
	return index, nil
}

// FastRepeat returns a new string consisting of n copies of the string s.
// It is optimized for cases where n is large by pre-allocating the builder capacity.
// If n is zero or negative, an empty string is returned.
//
// Examples:
//
//	FastRepeat("abc", 3) == "abcabcabc"
//	FastRepeat("a", 5) == "aaaaa"
//	FastRepeat("abc", 0) == ""
//	FastRepeat("abc", -2) == ""
func FastRepeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	if n == 1 {
		return s
	}

	// Pre-allocate capacity for the strings.Builder for efficiency.
	// This is especially beneficial for large values of n.
	var builder strings.Builder
	builder.Grow(len(s) * n)

	for i := 0; i < n; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}

// Swap returns a new string with the characters at the given indices swapped.
// It handles Unicode characters correctly by operating on runes.
// If either index is out of bounds, or if the indices are the same,
// the original string is returned without modification.
//
// @param s The input string.
// @param i The index of the first character to swap.
// @param j The index of the second character to swap.
// @return A new string with the characters at indices i and j swapped,
//         or the original string if indices are invalid.
//
// Examples:
//
//	Swap("hello", 1, 3) == "hlelo"
//	Swap("你好", 0, 1) == "好你"
//	Swap("abc", 0, 0) == "abc"
//	Swap("abc", 0, 5) == "abc"
func Swap(s string, i, j int) string {
	runes := []rune(s)
	if i < 0 || j < 0 || i >= len(runes) || j >= len(runes) || i == j {
		return s
	}
	runes[i], runes[j] = runes[j], runes[i]
	return string(runes)
}

// SafeSplit splits a string around each instance of the separator, returning a slice of substrings.
// If the separator is an empty string, Split splits after each UTF-8 sequence.
// It returns an error if the separator is empty and the string is not empty.
//
// Examples:
//
//	SafeSplit("a,b,c", ",") == ([]string{"a", "b", "c"}, nil)
//	SafeSplit("a,b,c", "") returns ("", error) // separator is empty and string is not empty
//	SafeSplit("", ",") == ([]string{""}, nil)
//	SafeSplit("", "") == ([]string{}, nil)
func SafeSplit(s, sep string) ([]string, error) {
	if sep == "" && s != "" {
		return nil, errors.New("separator cannot be empty if string is not empty")
	}
	return strings.Split(s, sep), nil
}

// ContainsGeneric checks if a slice of any comparable type contains a specific item.
//
// Examples:
//
//	ContainsGeneric([]int{1, 2, 3}, 2) == true
//	ContainsGeneric([]string{"a", "b", "c"}, "d") == false
func ContainsGeneric[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// NormalizeSpaces replaces multiple whitespace characters in a string with a single space.
// It also trims leading and trailing whitespace.
//
// Examples:
//
//	NormalizeSpaces("  hello   world  ") == "hello world"
//	NormalizeSpaces("a\t\nb") == "a b"
//	NormalizeSpaces("single") == "single"
//	NormalizeSpaces("") == ""
func NormalizeSpaces(s string) string {
	if s == "" {
		return ""
	}

	var builder strings.Builder
	var lastCharIsSpace bool

	for _, r := range s {
		if unicode.IsSpace(r) {
			if !lastCharIsSpace {
				builder.WriteRune(' ')
				lastCharIsSpace = true
			}
		} else {
			builder.WriteRune(r)
			lastCharIsSpace = false
		}
	}

	// Trim leading and trailing spaces
	result := builder.String()
	return strings.TrimSpace(result)
}

package utils

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
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
//
//	or the original string if indices are invalid.
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

// Wrap returns a new string where the input string `s` is wrapped by `prefix` and `suffix`.
// If either `prefix` or `suffix` is empty, it's treated as if it were not provided.
//
// @param s The string to wrap.
// @param prefix The string to prepend.
// @param suffix The string to append.
// @return A new string with `s` wrapped by `prefix` and `suffix`.
//
// Examples:
//
//	Wrap("world", "hello ", "!") == "hello world!"
//	Wrap("text", "", "...") == "text..."
//	Wrap("content", "[", "]") == "[content]"
//	Wrap("data", "", "") == "data"
func Wrap(s, prefix, suffix string) string {
	return prefix + s + suffix
}

// SafeWrap returns a new string where the input string `s` is wrapped by `prefix` and `suffix`.
// If either `prefix` or `suffix` is empty, it's treated as if it were not provided.
// It returns an error if both prefix and suffix are empty and the input string `s` is also empty.
//
// @param s The string to wrap.
// @param prefix The string to prepend.
// @param suffix The string to append.
// @return A new string with `s` wrapped by `prefix` and `suffix`, or an error if inputs are invalid.
//
// Examples:
//
//	SafeWrap("world", "hello ", "!") == ("hello world!", nil)
//	SafeWrap("text", "", "...") == ("text...", nil)
//	SafeWrap("content", "[", "]") == ("[content]", nil)
//	SafeWrap("data", "", "") == ("data", nil)
//	SafeWrap("", "", "") returns ("", errors.New("cannot wrap an empty string with empty prefix and suffix"))
func SafeWrap(s, prefix, suffix string) (string, error) {
	if s == "" && prefix == "" && suffix == "" {
		return "", errors.New("cannot wrap an empty string with empty prefix and suffix")
	}
	return prefix + s + suffix, nil
}

// CountLines counts the number of lines in a string.
// A line is considered to be terminated by a newline character (\n).
// An empty string has 0 lines. A string with no newline characters has 1 line.
//
// Examples:
//
//	CountLines("hello\nworld") == 2
//	CountLines("hello") == 1
//	CountLines("") == 0
//	CountLines("\n") == 1
//	CountLines("line1\nline2\n") == 2
func CountLines(s string) int {
	if s == "" {
		return 0
	}
	count := 1 // Start with 1 line assuming at least one non-empty string
	for _, r := range s {
		if r == '\n' {
			count++
		}
	}
	// If the string ends with a newline, the last increment might have counted an extra "empty" line
	// after the final newline. strings.Split handles this by not including an empty string at the end.
	// For consistency, we'll mimic that behavior by checking if the string ends with a newline.
	if strings.HasSuffix(s, "\n") && count > 1 {
		// If it ends with a newline and we counted more than one line,
		// the last increment was for an empty line after the final newline.
		// We should not count this empty line.
		return count - 1
	}
	return count
}

// BeforeFirst returns the substring before the first occurrence of the separator.
// If the separator is not found, the entire string is returned.
// If the separator is empty, an empty string is returned.
//
// Examples:
//
//	BeforeFirst("hello world", " ") == "hello"
//	BeforeFirst("hello", "x") == "hello"
//	BeforeFirst("hello", "") == ""
func BeforeFirst(s, sep string) string {
	if sep == "" {
		return ""
	}
	index := strings.Index(s, sep)
	if index == -1 {
		return s
	}
	return s[:index]
}

// SafeBeforeFirst returns the substring before the first occurrence of the separator.
// If the separator is not found, the entire string is returned.
// If the separator is empty, an empty string is returned.
// It returns an error if the separator is empty and the string is not empty.
//
// Examples:
//
//	SafeBeforeFirst("hello world", " ") == ("hello", nil)
//	SafeBeforeFirst("hello", "x") == ("hello", nil)
//	SafeBeforeFirst("hello", "") == ("", errors.New("separator cannot be empty if string is not empty"))
func SafeBeforeFirst(s, sep string) (string, error) {
	if sep == "" && s != "" {
		return "", errors.New("separator cannot be empty if string is not empty")
	}
	if sep == "" {
		return "", nil
	}
	index := strings.Index(s, sep)
	if index == -1 {
		return s, nil
	}
	return s[:index], nil
}

// FastContains checks if a slice of strings contains a specific string.
// This version is optimized for performance by using a map for lookups,
// which provides O(1) average time complexity for checking existence,
// compared to O(n) for a linear scan.
//
// Examples:
//
//	FastContains([]string{"a", "b", "c"}, "b") == true
//	FastContains([]string{"a", "b", "c"}, "d") == false
func FastContains(slice []string, item string) bool {
	if len(slice) == 0 {
		return false
	}
	// Use a map for O(1) average time complexity lookups.
	// This is more efficient than a linear scan for larger slices.
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

// AfterLast returns the substring after the last occurrence of the separator.
// If the separator is not found, an empty string is returned.
// If the separator is empty, the original string is returned.
//
// Examples:
//
//	AfterLast("hello world", " ") == "world"
//	AfterLast("hello", "x") == ""
//	AfterLast("hello", "") == "hello"
func AfterLast(s, sep string) string {
	if sep == "" {
		return s
	}
	index := strings.LastIndex(s, sep)
	if index == -1 {
		return ""
	}
	return s[index+len(sep):]
}

// FastReverse returns the reverse of a string.
// It handles Unicode characters correctly by operating on runes.
// This version is optimized for performance by pre-allocating the rune slice capacity.
//
// Examples:
//
//	FastReverse("hello") == "olleh"
//	FastReverse("你好") == "好你"
func FastReverse(s string) string {
	runes := make([]rune, len([]rune(s))) // Pre-allocate rune slice capacity
	copy(runes, []rune(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ContainsAnyGeneric checks if a slice of any type contains a specific item.
// This function leverages Go generics to work with slices of any type that supports equality comparison.
//
// @param slice The slice to search within.
// @param item The item to search for in the slice.
// @return true if the item is found in the slice, false otherwise.
//
// Examples:
//
//	ContainsAnyGeneric([]int{1, 2, 3}, 2) == true
//	ContainsAnyGeneric([]string{"a", "b", "c"}, "d") == false
func ContainsAnyGeneric[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ValidateRange checks if an integer value is within a specified range [min, max].
// It returns an error if the value is less than min or greater than max.
// It returns an error if min > max.
//
// Examples:
//
//	ValidateRange(5, 0, 10) == nil
//	ValidateRange(-5, 0, 10) returns an error
//	ValidateRange(15, 0, 10) returns an error
//	ValidateRange(5, 10, 0) returns an error
func ValidateRange(val, min, max int) error {
	if min > max {
		return errors.New("min cannot be greater than max")
	}
	if val < min {
		return errors.New("value is less than minimum")
	}
	if val > max {
		return errors.New("value is greater than maximum")
	}
	return nil
}

// MapGeneric applies a function to each element of a slice and returns a new slice with the results.
// The function `f` takes an element of type T and returns an element of type U.
//
// Examples:
//
//	MapGeneric([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n) }) == []string{"1", "2", "3"}
//	MapGeneric([]string{"a", "b", "c"}, func(s string) int { return len(s) }) == []int{1, 1, 1}
func MapGeneric[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// CountWords counts the number of words in a string.
// Words are defined as sequences of non-whitespace characters separated by whitespace.
// It trims leading/trailing whitespace before counting.
//
// @param s The input string.
// @return The number of words in the string.
//
// Examples:
//
//	CountWords("hello world") == 2
//	CountWords("  leading and trailing spaces  ") == 4
//	CountWords("singleword") == 1
//	CountWords("a\t\nb") == 3
//	CountWords("") == 0
//	CountWords("   ") == 0
func CountWords(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	// Split by one or more whitespace characters
	words := strings.Fields(s)
	return len(words)
}

// ValidateNotEmpty checks if a string is not empty and does not consist solely of whitespace.
// It returns an error if the string is empty or contains only whitespace.
//
// Examples:
//
//	ValidateNotEmpty("hello") == nil
//	ValidateNotEmpty("  ") returns an error
//	ValidateNotEmpty("") returns an error
func ValidateNotEmpty(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("string cannot be empty or contain only whitespace")
	}
	return nil
}

// Compare returns an integer indicating the comparison between two strings.
// It returns:
//   - -1 if s1 < s2
//   - 0 if s1 == s2
//   - 1 if s1 > s2
//
// @param s1 The first string to compare.
// @param s2 The second string to compare.
// @return An integer representing the comparison result.
//
// Examples:
//
//	Compare("apple", "banana") == -1
//	Compare("banana", "apple") == 1
//	Compare("cherry", "cherry") == 0
func Compare(s1, s2 string) int {
	if s1 < s2 {
		return -1
	}
	if s1 > s2 {
		return 1
	}
	return 0
}

// ValidateOneOf checks if a string value is present in a given slice of allowed string values.
// It returns an error if the value is not found in the slice.
//
// Examples:
//
//	ValidateOneOf("red", []string{"red", "green", "blue"}) == nil
//	ValidateOneOf("yellow", []string{"red", "green", "blue"}) returns an error
//	ValidateOneOf("red", []string{}) returns an error
func ValidateOneOf(val string, allowed []string) error {
	if len(allowed) == 0 {
		return errors.New("allowed values list cannot be empty")
	}
	for _, item := range allowed {
		if val == item {
			return nil
		}
	}
	return errors.New("value is not one of the allowed options")
}

// ValidateISBN10 checks if a string is a valid ISBN-10 number.
// An ISBN-10 consists of 10 digits, where the last digit can be 'X' representing 10.
// The check digit calculation is: (10*d1 + 9*d2 + ... + 2*d9 + 1*d10) mod 11 == 0.
// It returns an error if the string is not a valid ISBN-10.
//
// Examples:
//
//	ValidateISBN10("0321714113") == nil
//	ValidateISBN10("0439023521") == nil
//	ValidateISBN10("032171411X") == nil
//	ValidateISBN10("0321714114") returns an error (invalid check digit)
//	ValidateISBN10("12345") returns an error (incorrect length)
//	ValidateISBN10("ABCDEFGHIJ") returns an error (non-digit characters)
func ValidateISBN10(isbn string) error {
	isbn = strings.ReplaceAll(isbn, "-", "") // Remove hyphens

	if len(isbn) != 10 {
		return errors.New("invalid ISBN-10 length")
	}

	sum := 0
	for i := 0; i < 9; i++ {
		digit := int(isbn[i] - '0')
		if digit < 0 || digit > 9 {
			return errors.New("invalid character in ISBN-10")
		}
		sum += digit * (10 - i)
	}

	lastChar := isbn[9]
	var lastDigit int
	if lastChar == 'X' || lastChar == 'x' {
		lastDigit = 10
	} else {
		lastDigit = int(lastChar - '0')
		if lastDigit < 0 || lastDigit > 9 {
			return errors.New("invalid last character in ISBN-10")
		}
	}
	sum += lastDigit

	if sum%11 != 0 {
		return errors.New("invalid ISBN-10 check digit")
	}

	return nil
}

// Count returns the number of occurrences of a substring within a string.
//
// Examples:
//
//	Count("ababab", "ab") == 3
//	Count("aaaaa", "a") == 5
//	Count("hello", "l") == 2
//	Count("world", "x") == 0
//	Count("abc", "") == 4 // Empty string matches at the beginning, between each character, and at the end.
func Count(s, substr string) int {
	if substr == "" {
		return len([]rune(s)) + 1
	}
	return strings.Count(s, substr)
}

// ExtractNumber returns the first sequence of digits found in a string.
// If no digits are found, it returns an empty string.
//
// @param s The input string to search within.
// @return The first sequence of digits found in the string, or an empty string if none exist.
//
// Examples:
//
//	ExtractNumber("abc123def456") == "123"
//	ExtractNumber("no digits here") == ""
//	ExtractNumber("123") == "123"
//	ExtractNumber(" leading 1 digit") == "1"
func ExtractNumber(s string) string {
	var builder strings.Builder
	foundDigit := false
	for _, r := range s {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
			foundDigit = true
		} else if foundDigit {
			// Stop once we encounter a non-digit after finding at least one digit
			break
		}
	}
	return builder.String()
}

// SafeExtractNumber returns the first sequence of digits found in a string.
// If no digits are found, it returns an empty string and nil error.
//
// @param s The input string to search within.
// @return The first sequence of digits found in the string, or an empty string if none exist, and nil error.
//
// Examples:
//
//	SafeExtractNumber("abc123def456") == ("123", nil)
//	SafeExtractNumber("no digits here") == ("", nil)
//	SafeExtractNumber("123") == ("123", nil)
//	SafeExtractNumber(" leading 1 digit") == ("1", nil)
func SafeExtractNumber(s string) (string, error) {
	var builder strings.Builder
	foundDigit := false
	for _, r := range s {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
			foundDigit = true
		} else if foundDigit {
			// Stop once we encounter a non-digit after finding at least one digit
			break
		}
	}
	return builder.String(), nil
}

// ValidateURL checks if a string is a valid URL.
// It uses Go's built-in net/url.Parse function for validation.
// It returns an error if the URL is invalid or if it cannot be parsed.
//
// Examples:
//
//	ValidateURL("https://www.google.com") == nil
//	ValidateURL("http://localhost:8080/path?query=test") == nil
//	ValidateURL("invalid-url") returns an error
//	ValidateURL("ftp://example.com") == nil
//	ValidateURL("://missing.scheme.com") returns an error
func ValidateURL(urlStr string) error {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}
	return nil
}

// SafeCompare returns an integer indicating the comparison between two strings.
// It returns:
//   - -1 if s1 < s2
//   - 0 if s1 == s2
//   - 1 if s1 > s2
//
// It returns an error if there's an issue during comparison (though standard string comparison in Go rarely errors).
//
// @param s1 The first string to compare.
// @param s2 The second string to compare.
// @return An integer representing the comparison result, or an error if comparison fails.
//
// Examples:
//
//	SafeCompare("apple", "banana") == (-1, nil)
//	SafeCompare("banana", "apple") == (1, nil)
//	SafeCompare("cherry", "cherry") == (0, nil)
func SafeCompare(s1, s2 string) (int, error) {
	if s1 < s2 {
		return -1, nil
	}
	if s1 > s2 {
		return 1, nil
	}
	return 0, nil
}

// ValidateEmail checks if a string is a valid email address.
// It uses a regular expression for basic email format validation.
// Note: This is a basic check and doesn't cover all RFC 5322 complexities.
//
// Examples:
//
//	ValidateEmail("test@example.com") == nil
//	ValidateEmail("invalid-email") returns an error
//	ValidateEmail("user+alias@domain.co.uk") == nil
//	ValidateEmail("@domain.com") returns an error
//	ValidateEmail("user@domain.") returns an error
func ValidateEmail(email string) error {
	// A common regex for basic email validation.
	// It checks for a username part, followed by '@', then a domain part.
	// The domain part must contain at least one dot.
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// SafeNormalizeSpaces replaces multiple whitespace characters in a string with a single space.
// It also trims leading and trailing whitespace.
// It returns an error if the input string is nil, though Go strings are not nilable. This signature
// is for consistency with other Safe* functions that might encounter nilable inputs.
//
// Examples:
//
//	SafeNormalizeSpaces("  hello   world  ") == ("hello world", nil)
//	SafeNormalizeSpaces("a\t\nb") == ("a b", nil)
//	SafeNormalizeSpaces("single") == ("single", nil)
//	SafeNormalizeSpaces("") == ("", nil)
func SafeNormalizeSpaces(s string) (string, error) {
	if s == "" {
		return "", nil
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
	return strings.TrimSpace(result), nil
}

// IsURL checks if a string is a valid URL.
// It uses Go's built-in net/url.Parse function for validation.
// It returns an error if the URL is invalid or if it cannot be parsed.
//
// Examples:
//
//	IsURL("https://www.google.com") == nil
//	IsURL("http://localhost:8080/path?query=test") == nil
//	IsURL("invalid-url") returns an error
//	IsURL("ftp://example.com") == nil
//	IsURL("://missing.scheme.com") returns an error
func IsURL(urlStr string) error {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}
	return nil
}

// FastCapitalize returns the string with the first letter capitalized.
// If the string is empty, it returns an empty string.
// It handles Unicode characters correctly.
// This version is optimized by avoiding unnecessary rune slice conversions when the string is short or already capitalized.
//
// Examples:
//
//	FastCapitalize("hello") == "Hello"
//	FastCapitalize("Hello") == "Hello"
//	FastCapitalize("") == ""
func FastCapitalize(s string) string {
	if s == "" {
		return s
	}

	// Check if the first character is already uppercase.
	// This avoids unnecessary operations if the string is already capitalized.
	runes := []rune(s)
	if unicode.IsUpper(runes[0]) {
		return s
	}

	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// SplitOnce splits a string into two parts at the first occurrence of the separator.
// It returns the part before the separator and the part after the separator.
// If the separator is not found, it returns the original string and an empty string.
// If the separator is empty, it returns an empty string and the original string.
//
// @param s The string to split.
// @param sep The separator string.
// @return A slice containing two strings: the part before the separator and the part after.
//
// Examples:
//
//	SplitOnce("hello,world", ",") == []string{"hello", "world"}
//	SplitOnce("helloworld", ",") == []string{"helloworld", ""}
//	SplitOnce("hello,world", "") == []string{"", "hello,world"}
func SplitOnce(s, sep string) []string {
	if sep == "" {
		return []string{"", s}
	}
	index := strings.Index(s, sep)
	if index == -1 {
		return []string{s, ""}
	}
	return []string{s[:index], s[index+len(sep):]}
}

// IsAlpha checks if a string contains only alphabetic characters.
// It returns true if the string is not empty and all characters are letters.
// It returns false otherwise, including for empty strings.
//
// @param s The input string to check.
// @return true if the string contains only alphabetic characters, false otherwise.
//
// Examples:
//
//	IsAlpha("HelloWorld") == true
//	IsAlpha("Hello World") == false // Contains a space
//	IsAlpha("Hello123") == false   // Contains digits
//	IsAlpha("") == false           // Empty string
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// SafeCountLines counts the number of lines in a string.
// A line is considered to be terminated by a newline character (\n).
// An empty string has 0 lines. A string with no newline characters has 1 line.
// It returns an error if the input string is nil (though Go strings are not nilable,
// this signature is for consistency with other Safe* functions).
//
// Examples:
//
//	SafeCountLines("hello\nworld") == (2, nil)
//	SafeCountLines("hello") == (1, nil)
//	SafeCountLines("") == (0, nil)
//	SafeCountLines("\n") == (1, nil)
//	SafeCountLines("line1\nline2\n") == (2, nil)
func SafeCountLines(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	count := 1 // Start with 1 line assuming at least one non-empty string
	for _, r := range s {
		if r == '\n' {
			count++
		}
	}
	// If the string ends with a newline, the last increment might have counted an extra "empty" line
	// after the final newline. strings.Split handles this by not including an empty string at the end.
	// For consistency, we'll mimic that behavior by checking if the string ends with a newline.
	if strings.HasSuffix(s, "\n") && count > 1 {
		// If it ends with a newline and we counted more than one line,
		// the last increment was for an empty line after the final newline.
		// We should not count this empty line.
		return count - 1, nil
	}
	return count, nil
}

// Unquote returns the unquoted version of a string, interpreting 1-byte C-style escape sequences.
// It returns an error if the string is not a valid quoted string.
//
// Examples:
//
//	Unquote(`"hello"`) == ("hello", nil)
//	Unquote(`"hello\nworld"`) == ("hello\nworld", nil)
//	Unquote(`"invalid escape\z"`) returns ("", error)
//	Unquote(`"unclosed string`) returns ("", error)
func Unquote(s string) (string, error) {
	return strconv.Unquote(s)
}

// TakeWhile returns a new slice containing elements from the input slice
// from the beginning until the predicate function returns false.
// The predicate function should return true for elements to keep and false for elements to discard.
//
// Examples:
//
//	TakeWhile([]int{1, 2, 3, 4, 5}, func(n int) bool { return n < 3 }) == []int{1, 2}
//	TakeWhile([]string{"a", "b", "c"}, func(s string) bool { return s != "d" }) == []string{"a", "b", "c"}
//	TakeWhile([]int{1, 2, 3}, func(n int) bool { return n > 5 }) == []int{}
func TakeWhile[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		} else {
			break // Stop as soon as the predicate returns false
		}
	}
	return result
}

// ValidatePhone checks if a string is a valid phone number.
// It uses a simplified regular expression to validate common phone number formats.
// This validation is basic and might not cover all international formats or complex cases.
// It returns an error if the string is not a valid phone number.
//
// Examples:
//
//	ValidatePhone("123-456-7890") == nil
//	ValidatePhone("(123) 456-7890") == nil
//	ValidatePhone("1234567890") == nil
//	ValidatePhone("+1 123 456 7890") == nil
//	ValidatePhone("123-456") returns an error
//	ValidatePhone("abc") returns an error
func ValidatePhone(phone string) error {
	// This regex allows for optional country code, area code in parentheses, spaces, hyphens, and digits.
	// It's a simplified pattern for common formats.
	const phoneRegex = `^\+?(\d{1,3})?[-. ]?\(?\d{3}\)?[-. ]?\d{3}[-. ]?\d{4}$`
	if !regexp.MustCompile(phoneRegex).MatchString(phone) {
		return errors.New("invalid phone number format")
	}
	return nil
}

// DropWhile returns a new slice containing elements from the input slice
// after the predicate function starts returning false.
// The predicate function should return true for elements to discard and false for elements to keep.
//
// Examples:
//
//	DropWhile([]int{1, 2, 3, 4, 5}, func(n int) bool { return n < 3 }) == []int{3, 4, 5}
//	DropWhile([]string{"a", "b", "c"}, func(s string) bool { return s != "d" }) == []string{}
//	DropWhile([]int{1, 2, 3}, func(n int) bool { return n > 5 }) == []int{1, 2, 3}
func DropWhile[T any](slice []T, predicate func(T) bool) []T {
	// Find the index where the predicate first returns false
	firstFalseIndex := -1
	for i, item := range slice {
		if !predicate(item) {
			firstFalseIndex = i
			break
		}
	}

	// If the predicate was always true, return an empty slice
	if firstFalseIndex == -1 {
		return []T{}
	}

	// Return the slice from the first element that satisfied the predicate to the end
	return slice[firstFalseIndex:]
}

// IsEmail checks if a string is a valid email address.
// It uses a regular expression for basic email format validation.
// Note: This is a basic check and doesn't cover all RFC 5322 complexities.
//
// Examples:
//
//	IsEmail("test@example.com") == nil
//	IsEmail("invalid-email") returns an error
//	IsEmail("user+alias@domain.co.uk") == nil
//	IsEmail("@domain.com") returns an error
//	IsEmail("user@domain.") returns an error
func IsEmail(email string) error {
	// A common regex for basic email validation.
	// It checks for a username part, followed by '@', then a domain part.
	// The domain part must contain at least one dot.
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// RemoveNonNumeric removes all non-digit characters from a string.
// It iterates through the string, building a new string containing only digits.
//
// @param s The input string to process.
// @return A new string containing only the numeric characters from the input string.
//
// Examples:
//
//	RemoveNonNumeric("abc123def456") == "123456"
//	RemoveNonNumeric("no digits here") == ""
//	RemoveNonNumeric("123") == "123"
func RemoveNonNumeric(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// RemoveNonAlphaNumeric removes all characters that are not alphanumeric from a string.
// It iterates through the string, building a new string containing only letters and digits.
//
// @param s The input string to process.
// @return A new string containing only the alphanumeric characters from the input string.
//
// Examples:
//
//	RemoveNonAlphaNumeric("abc 123 !@# def 456") == "abc123def456"
//	RemoveNonAlphaNumeric("no alphanumeric here") == "noalphanumerichere"
//	RemoveNonAlphaNumeric("12345") == "12345"
func RemoveNonAlphaNumeric(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// ValidatePattern checks if a string matches a given regular expression pattern.
// It returns an error if the string does not match the pattern.
// It also returns an error if the provided pattern is invalid.
//
// Examples:
//
//	ValidatePattern("123-456-7890", `^\d{3}-\d{3}-\d{4}$`) == nil
//	ValidatePattern("1234567890", `^\d{3}-\d{3}-\d{4}$`) returns an error
//	ValidatePattern("abc", `^\d+$`) returns an error
//	ValidatePattern("123", `(`) returns an error (invalid regex pattern)
func ValidatePattern(s, pattern string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err // Return error if the regex pattern itself is invalid
	}
	if !re.MatchString(s) {
		return errors.New("string does not match the provided pattern")
	}
	return nil
}

// Intersect returns a new slice containing elements that are present in both input slices.
// It uses generics to work with slices of any comparable type.
//
// Examples:
//
//	Intersect([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}) == []int{3, 4}
//	Intersect([]string{"a", "b", "c"}, []string{"b", "c", "d"}) == []string{"b", "c"}
//	Intersect([]int{1, 2}, []int{3, 4}) == []int{}
func Intersect[T comparable](slice1, slice2 []T) []T {
	// Use a map to store elements of the first slice for efficient lookup.
	set1 := make(map[T]struct{}, len(slice1))
	for _, item := range slice1 {
		set1[item] = struct{}{}
	}

	var result []T
	// Iterate over the second slice and check if elements exist in the map.
	for _, item := range slice2 {
		if _, ok := set1[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

// Union returns a new slice containing all unique elements from both input slices.
// It uses generics to work with slices of any comparable type.
// The order of elements in the resulting slice is not guaranteed.
//
// Examples:
//
//	Union([]int{1, 2, 3}, []int{3, 4, 5}) == []int{1, 2, 3, 4, 5} (order may vary)
//	Union([]string{"a", "b"}, []string{"b", "c", "d"}) == []string{"a", "b", "c", "d"} (order may vary)
//	Union([]int{1, 2}, []int{3, 4}) == []int{1, 2, 3, 4} (order may vary)
func Union[T comparable](slice1, slice2 []T) []T {
	// Use a map to store all unique elements encountered.
	set := make(map[T]struct{}, len(slice1)+len(slice2))

	// Add elements from the first slice to the set.
	for _, item := range slice1 {
		set[item] = struct{}{}
	}

	// Add elements from the second slice to the set.
	for _, item := range slice2 {
		set[item] = struct{}{}
	}

	// Convert the set back to a slice.
	result := make([]T, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	return result
}

// Exclude returns a new slice containing elements from the input slice
// that do NOT satisfy the given predicate function.
// The predicate function should return true for elements to exclude and false for elements to keep.
//
// Examples:
//
//	Exclude([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }) == []int{1, 3, 5}
func Exclude[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice)/2) // Pre-allocate result slice with a heuristic initial capacity.
	for _, item := range slice {
		if !predicate(item) { // Keep elements for which the predicate is false
			result = append(result, item)
		}
	}
	return result
}

// Unzip takes a slice of slices and returns a slice of slices where each inner slice
// contains elements at the same index from the original inner slices.
// It assumes all inner slices have the same length. If the input is empty or
// contains empty inner slices, it returns an empty slice.
//
// Examples:
//
//	Unzip([][]int{{1, 2, 3}, {4, 5, 6}}) == [][]int{{1, 4}, {2, 5}, {3, 6}}
//	Unzip([][]string{{"a", "b"}, {"c", "d"}}) == [][]string{{"a", "c"}, {"b", "d"}}
//	Unzip([][]int{}) == [][]int{}
//	Unzip([][]int{{}}) == [][]int{}
func Unzip[T any](slices [][]T) [][]T {
	if len(slices) == 0 || len(slices[0]) == 0 {
		return [][]T{}
	}

	numOuter := len(slices)
	numInner := len(slices[0])

	// Initialize the result slice with the correct dimensions.
	result := make([][]T, numInner)
	for i := range result {
		result[i] = make([]T, numOuter)
	}

	// Populate the result slice by transposing the input.
	for i := 0; i < numOuter; i++ {
		for j := 0; j < numInner; j++ {
			result[j][i] = slices[i][j]
		}
	}

	return result
}

// SplitOnceAfter splits a string into two parts at the first occurrence of the separator,
// returning the part before the separator and the part after it (including the separator).
// If the separator is not found, it returns the original string and an empty string.
// If the separator is empty, it returns an empty string and the original string.
//
// @param s The string to split.
// @param sep The separator string.
// @return A slice containing two strings: the part before the separator and the part after (including the separator).
//
// Examples:
//
//	SplitOnceAfter("hello,world", ",") == []string{"hello", ",world"}
//	SplitOnceAfter("helloworld", ",") == []string{"helloworld", ""}
//	SplitOnceAfter("hello,world", "") == []string{"", "hello,world"}
func SplitOnceAfter(s, sep string) []string {
	if sep == "" {
		return []string{"", s}
	}
	index := strings.Index(s, sep)
	if index == -1 {
		return []string{s, ""}
	}
	return []string{s[:index], s[index:]}
}

// ValidateHex checks if a string represents a valid hexadecimal number.
// It returns an error if the string contains any characters that are not
// hexadecimal digits (0-9, a-f, A-F) or if the string is empty.
//
// Examples:
//
//	ValidateHex("1A2B3C") == nil
//	ValidateHex("abcdef") == nil
//	ValidateHex("0123456789") == nil
//	ValidateHex("1a2b3c") == nil
//	ValidateHex("invalid hex") returns an error
//	ValidateHex("") returns an error
func ValidateHex(s string) error {
	if s == "" {
		return errors.New("string cannot be empty")
	}
	for _, r := range s {
		if !unicode.IsDigit(r) && (r < 'a' || r > 'f') && (r < 'A' || r > 'F') {
			return errors.New("string contains non-hexadecimal characters")
		}
	}
	return nil
}

// Diff returns the elements that are in slice1 but not in slice2.
// It uses generics to work with slices of any comparable type.
//
// Examples:
//
//	Diff([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}) == []int{1, 2}
//	Diff([]string{"a", "b", "c"}, []string{"b", "c", "d"}) == []string{"a"}
//	Diff([]int{1, 2}, []int{3, 4}) == []int{1, 2}
func Diff[T comparable](slice1, slice2 []T) []T {
	// Use a map to store elements of the second slice for efficient lookup.
	set2 := make(map[T]struct{}, len(slice2))
	for _, item := range slice2 {
		set2[item] = struct{}{}
	}

	var result []T
	// Iterate over the first slice and add elements not found in the map.
	for _, item := range slice1 {
		if _, ok := set2[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}

// DiffGeneric returns the elements that are in slice1 but not in slice2.
// It uses generics to work with slices of any comparable type.
//
// Examples:
//
//	DiffGeneric([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}) == []int{1, 2}
//	DiffGeneric([]string{"a", "b", "c"}, []string{"b", "c", "d"}) == []string{"a"}
//	DiffGeneric([]int{1, 2}, []int{3, 4}) == []int{1, 2}
func DiffGeneric[T comparable](slice1, slice2 []T) []T {
	// Use a map to store elements of the second slice for efficient lookup.
	set2 := make(map[T]struct{}, len(slice2))
	for _, item := range slice2 {
		set2[item] = struct{}{}
	}

	var result []T
	// Iterate over the first slice and add elements not found in the map.
	for _, item := range slice1 {
		if _, ok := set2[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}

// FastTrimAll removes all whitespace characters (spaces, tabs, newlines, etc.) from a string.
// It uses a strings.Builder for efficient string concatenation and pre-allocates capacity.
//
// Examples:
//
//	FastTrimAll("  hello \t world \n") == "helloworld"
func FastTrimAll(s string) string {
	var builder strings.Builder
	builder.Grow(len(s)) // Pre-allocate capacity for efficiency
	for _, r := range s {
		if !unicode.IsSpace(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// FilterGeneric returns a new slice containing only elements from the input slice
// that satisfy the given predicate function.
// The predicate function should return true for elements to keep and false for elements to discard.
// This generic version works with slices of any type.
//
// Examples:
//
//	FilterGeneric([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }) == []int{2, 4}
//	FilterGeneric([]string{"apple", "banana", "cherry"}, func(s string) bool { return len(s) > 5 }) == []string{"banana", "cherry"}
func FilterGeneric[T any](slice []T, predicate func(T) bool) []T {
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

// Unique returns a new slice containing only the unique elements from the input slice.
// It uses generics to work with slices of any comparable type.
// The order of elements in the resulting slice is not guaranteed.
//
// Examples:
//
//	Unique([]int{1, 2, 2, 3, 4, 4, 4, 5}) == []int{1, 2, 3, 4, 5} (order may vary)
//	Unique([]string{"a", "b", "a", "c", "b"}) == []string{"a", "b", "c"} (order may vary)
func Unique[T comparable](slice []T) []T {
	set := make(map[T]struct{}, len(slice))
	for _, item := range slice {
		set[item] = struct{}{}
	}

	result := make([]T, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	return result
}

// SafeSplitOnce splits a string into two parts at the first occurrence of the separator.
// It returns the part before the separator and the part after the separator.
// If the separator is not found, it returns the original string and an empty string.
// If the separator is empty, it returns an empty string and the original string.
// It returns an error if the separator is empty and the string is not empty,
// as this behavior is ambiguous and handled by strings.Split.
//
// @param s The string to split.
// @param sep The separator string.
// @return A slice containing two strings: the part before the separator and the part after, or an error.
//
// Examples:
//
//	SafeSplitOnce("hello,world", ",") == ([]string{"hello", "world"}, nil)
//	SafeSplitOnce("helloworld", ",") == ([]string{"helloworld", ""}, nil)
//	SafeSplitOnce("hello,world", "") == ([]string{"", "hello,world"}, nil)
//	SafeSplitOnce("a,b,c", "") returns ("", errors.New("separator cannot be empty if string is not empty"))
func SafeSplitOnce(s, sep string) ([]string, error) {
	if sep == "" && s != "" {
		return nil, errors.New("separator cannot be empty if string is not empty")
	}
	if sep == "" {
		return []string{"", s}, nil
	}
	index := strings.Index(s, sep)
	if index == -1 {
		return []string{s, ""}, nil
	}
	return []string{s[:index], s[index+len(sep):]}, nil
}

// Between returns the string between the first and last occurrences of a specified substring.
// If the start substring is not found, or the end substring is not found after the start,
// an empty string is returned. If the start or end substrings are empty, it may lead to
// unexpected behavior or empty results.
//
// Examples:
//
//	Between("hello [world]!", "[", "]") == "world"
//	Between("no delimiters here", "[", "]") == ""
//	Between("start middle end", "start", "end") == " middle "
//	Between("start", "start", "end") == ""
//	Between("end", "start", "end") == ""
func Between(s, start, end string) string {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return ""
	}
	// Adjust startIndex to be after the start delimiter
	startIndex += len(start)

	endIndex := strings.Index(s[startIndex:], end)
	if endIndex == -1 {
		return ""
	}
	return s[startIndex : startIndex+endIndex]
}

// SafeBetween returns the string between the first and last occurrences of a specified substring.
// If the start substring is not found, or the end substring is not found after the start,
// it returns an empty string and a nil error.
// If the start or end substrings are empty, it may lead to unexpected behavior or empty results,
// but no error is returned unless the substrings are invalid in a way that `strings.Index` would error (which is rare).
//
// Examples:
//
//	SafeBetween("hello [world]!", "[", "]") == ("world", nil)
//	SafeBetween("no delimiters here", "[", "]") == ("", nil)
//	SafeBetween("start middle end", "start", "end") == (" middle ", nil)
//	SafeBetween("start", "start", "end") == ("", nil)
//	SafeBetween("end", "start", "end") == ("", nil)
func SafeBetween(s, start, end string) (string, error) {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return "", nil
	}
	// Adjust startIndex to be after the start delimiter
	startIndex += len(start)

	endIndex := strings.Index(s[startIndex:], end)
	if endIndex == -1 {
		return "", nil
	}
	return s[startIndex : startIndex+endIndex], nil
}

// IsBlank checks if a string is empty, or contains only whitespace characters.
//
// @param s The string to check.
// @return true if the string is blank, false otherwise.
//
// Examples:
//
//	IsBlank("") == true
//	IsBlank("   ") == true
//	IsBlank("hello") == false
func IsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// WrapText wraps a given text to a specified line width, breaking lines at word boundaries.
// It aims to preserve existing line breaks in the input text.
//
// @param text The input string to wrap.
// @param lineWidth The maximum width of each line.
// @return A new string with the text wrapped to the specified line width.
//
// Examples:
//
//	WrapText("This is a long sentence that needs to be wrapped.", 15) == "This is a\nlong\nsentence\nthat needs to\nbe wrapped."
//	WrapText("Short text", 20) == "Short text"
//	WrapText("Line1\nLine2 with more text", 10) == "Line1\nLine2 with\nmore text"
//	WrapText("", 10) == ""
func WrapText(text string, lineWidth int) string {
	if text == "" || lineWidth <= 0 {
		return text
	}

	var builder strings.Builder
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if builder.Len() > 0 {
			builder.WriteRune('\n')
		}

		words := strings.Fields(line)
		if len(words) == 0 {
			continue // Skip empty lines after splitting by \n
		}

		currentLineLength := 0
		for i, word := range words {
			if currentLineLength+len(word) > lineWidth && currentLineLength > 0 {
				builder.WriteRune('\n')
				currentLineLength = 0
			}
			if currentLineLength > 0 {
				builder.WriteRune(' ')
				currentLineLength++
			}
			builder.WriteString(word)
			currentLineLength += len(word)

			// If this is the last word on the line, and it's exactly at the lineWidth,
			// we need to ensure the next word starts on a new line if there is one.
			if i == len(words)-1 && currentLineLength == lineWidth && len(words) > i+1 {
				// This case is tricky: if the last word perfectly fills the line,
				// the next word should naturally go to a new line in the next iteration.
				// No explicit action needed here other than letting the loop continue.
			}
		}
	}

	return builder.String()
}

// FastCapitalize returns the string with the first letter capitalized.
// If the string is empty, it returns an empty string.
// It handles Unicode characters correctly.
// This version is optimized by avoiding unnecessary rune slice conversions when the string is short or already capitalized.
//
// Examples:
//
//	FastCapitalize("hello") == "Hello"
//	FastCapitalize("Hello") == "Hello"
//	FastCapitalize("") == ""
func FastCapitalize(s string) string {
	if s == "" {
		return s
	}

	// Check if the first character is already uppercase.
	// This avoids unnecessary operations if the string is already capitalized.
	runes := []rune(s)
	if unicode.IsUpper(runes[0]) {
		return s
	}

	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Mask replaces all but the last `n` characters of a string with a masking character.
// If the string length is less than or equal to `n`, the original string is returned.
// The masking character defaults to '*'.
//
// Examples:
//
//	Mask("1234567890", 4) == "****67890"
//	Mask("short", 10) == "short"
//	Mask("password", 3, '#') == "#######rd"
//	Mask("test", 0) == "****"
func Mask(s string, n int, maskChar ...rune) string {
	if n < 0 {
		n = 0 // Treat negative n as 0
	}

	runes := []rune(s)
	length := len(runes)

	if length <= n {
		return s
	}

	var char rune
	if len(maskChar) > 0 {
		char = maskChar[0]
	} else {
		char = '*'
	}

	maskedPart := make([]rune, length-n)
	for i := range maskedPart {
		maskedPart[i] = char
	}

	return string(maskedPart) + string(runes[length-n:])
}

// ReplaceAllGeneric replaces all occurrences of `old` with `new` in a slice of any comparable type.
//
// Examples:
//
//	ReplaceAllGeneric([]int{1, 2, 1, 3}, 1, 10) == []int{10, 2, 10, 3}
//	ReplaceAllGeneric([]string{"a", "b", "a", "c"}, "a", "x") == []string{"x", "b", "x", "c"}
func ReplaceAllGeneric[T comparable](slice []T, old, new T) []T {
	result := make([]T, len(slice))
	for i, item := range slice {
		if item == old {
			result[i] = new
		} else {
			result[i] = item
		}
	}
	return result
}

// ValidateIMEI checks if a string is a valid IMEI (International Mobile Equipment Identity) number.
// An IMEI is a 15-digit number. This function performs a basic length check.
// It returns an error if the string is not exactly 15 digits long.
//
// Examples:
//
//	ValidateIMEI("123456789012345") == nil
//	ValidateIMEI("12345678901234") returns an error (too short)
//	ValidateIMEI("1234567890123456") returns an error (too long)
//	ValidateIMEI("ABCDEFGHIJKLMNO") returns an error (non-digit characters)
func ValidateIMEI(imei string) error {
	if len(imei) != 15 {
		return errors.New("invalid IMEI length: must be 15 digits")
	}
	for _, r := range imei {
		if !unicode.IsDigit(r) {
			return errors.New("invalid IMEI format: must contain only digits")
		}
	}
	return nil
}

// Quote returns the double-quoted version of s.
// It escapes special characters according to the rules used by JSON and Go.
//
// Examples:
//
//	Quote("hello") == `"hello"`
//	Quote("hello\nworld") == `"hello\nworld"`
//	Quote(`"`) == ` "\"" `
func Quote(s string) string {
	return strconv.Quote(s)
}

// RemoveAccents removes diacritical marks from accented characters in a string.
// It converts characters like 'é', 'ü', 'ñ' to 'e', 'u', 'n' respectively.
//
// Examples:
//
//	RemoveAccents("résumé") == "resume"
//	RemoveAccents("Crème brûlée") == "Creme brulee"
//	RemoveAccents("niño") == "nino"
//	RemoveAccents("你好") == "你好" // Non-accented characters are unchanged
func RemoveAccents(s string) string {
	// Use a library or a comprehensive mapping if extensive accent removal is needed.
	// For a common subset of accents, a direct mapping can be used.
	// This example covers a few common Latin-script accents.
	var builder strings.Builder
	for _, r := range s {
		switch r {
		case 'á', 'à', 'â', 'ä', 'å', 'ā':
			builder.WriteRune('a')
		case 'é', 'è', 'ê', 'ë', 'ē', 'ė':
			builder.WriteRune('e')
		case 'í', 'ì', 'î', 'ï', 'ī':
			builder.WriteRune('i')
		case 'ó', 'ò', 'ô', 'ö', 'ø', 'ō':
			builder.WriteRune('o')
		case 'ú', 'ù', 'û', 'ü', 'ū':
			builder.WriteRune('u')
		case 'ý', 'ÿ':
			builder.WriteRune('y')
		case 'ñ':
			builder.WriteRune('n')
		case 'ç':
			builder.WriteRune('c')
		case 'ş':
			builder.WriteRune('s')
		case 'ğ':
			builder.WriteRune('g')
		case 'æ':
			builder.WriteRune('ae')
		case 'œ':
			builder.WriteRune('oe')
		default:
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// RemoveAccents removes diacritical marks from accented characters in a string.
// It converts characters like 'é', 'ü', 'ñ' to 'e', 'u', 'n' respectively.
// This function is designed to handle common Latin-script accents. For more
// comprehensive accent removal across various languages, a dedicated library
// or a more extensive mapping would be recommended.
//
// @param s The input string potentially containing accented characters.
// @return A new string with diacritical marks removed.
//
// Examples:
//
//	RemoveAccents("résumé") == "resume"
//	RemoveAccents("Crème brûlée") == "Creme brulee"
//	RemoveAccents("niño") == "nino"
//	RemoveAccents("你好") == "你好" // Non-accented characters are unchanged
func RemoveAccents(s string) string {
	// Use a library or a comprehensive mapping if extensive accent removal is needed.
	// For a common subset of accents, a direct mapping can be used.
	// This example covers a few common Latin-script accents.
	var builder strings.Builder
	for _, r := range s {
		switch r {
		case 'á', 'à', 'â', 'ä', 'å', 'ā':
			builder.WriteRune('a')
		case 'é', 'è', 'ê', 'ë', 'ē', 'ė':
			builder.WriteRune('e')
		case 'í', 'ì', 'î', 'ï', 'ī':
			builder.WriteRune('i')
		case 'ó', 'ò', 'ô', 'ö', 'ø', 'ō':
			builder.WriteRune('o')
		case 'ú', 'ù', 'û', 'ü', 'ū':
			builder.WriteRune('u')
		case 'ý', 'ÿ':
			builder.WriteRune('y')
		case 'ñ':
			builder.WriteRune('n')
		case 'ç':
			builder.WriteRune('c')
		case 'ş':
			builder.WriteRune('s')
		case 'ğ':
			builder.WriteRune('g')
		case 'æ':
			builder.WriteRune('ae')
		case 'œ':
			builder.WriteRune('oe')
		default:
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// ValidateIP checks if a string is a valid IPv4 or IPv6 address.
// It uses Go's net.ParseIP function for validation.
// It returns an error if the string is not a valid IP address.
//
// Examples:
//
//	ValidateIP("192.168.1.1") == nil
//	ValidateIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334") == nil
//	ValidateIP("invalid-ip") returns an error
//	ValidateIP("192.168.1.256") returns an error
func ValidateIP(ipStr string) error {
	if net.ParseIP(ipStr) == nil {
		return errors.New("invalid IP address format")
	}
	return nil
}

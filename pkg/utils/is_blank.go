package utils

import (
	"unicode"
)

// IsBlank checks whether a string is empty or consists entirely of whitespace.
func IsBlank(s string) bool {
	if s == "" {
		return true
	}

	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

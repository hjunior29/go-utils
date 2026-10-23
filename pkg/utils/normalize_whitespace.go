package utils

import (
	"strings"
)

// NormalizeWhitespace trims leading and trailing whitespace and collapses
// any sequence of consecutive whitespace characters into a single ASCII space.
func NormalizeWhitespace(s string) string {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return ""
	}
	return strings.Join(fields, " ")
}

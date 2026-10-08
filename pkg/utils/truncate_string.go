package utils

// TruncateString truncates a string to a nonnegative maximum number of Unicode code points, without an ellipsis.
func TruncateString(s string, maxRunes int) string {
	if maxRunes < 0 {
		return s
	}

	runes := []rune(s)
	if len(runes) <= maxRunes {
		return s
	}

	return string(runes[:maxRunes])
}

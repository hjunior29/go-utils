package utils

import "unicode/utf8"

func CommonPrefix(a, b string) string {
	var lastMatch int
	for i, rA := range a {
		if i >= len(b) {
			break
		}
		rB, size := utf8.DecodeRuneInString(b[i:])
		if rA != rB {
			break
		}
		lastMatch = i + size
	}
	return a[:lastMatch]
}

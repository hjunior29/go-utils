package utils

import (
	"errors"
	"unicode/utf8"
)

// HammingDistance counts the differing Unicode code points between two equal-length strings.
// It returns an error if the strings have different lengths measured in bytes, or if they
// contain differing numbers of Unicode code points.
func HammingDistance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("strings must be of equal length")
	}

	runesA := []rune(a)
	runesB := []rune(b)
	distance := 0

	for i := 0; i < len(runesA); i++ {
		if runesA[i] != runesB[i] {
			distance++
		}
	}

	return distance, nil
}

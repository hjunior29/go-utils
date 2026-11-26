package utils

import (
	"errors"
	"strings"
)

func MaskString(s string, n int) (string, error) {
	if n < 0 {
		return "", errors.New("negative N is not allowed")
	}

	runnes := []rune(s)
	length := len(runnes)

	if n >= length {
		return s, nil
	}

	var builder strings.Builder
	builder.Grow(length)

	maskedCount := length - n
	for i := 0; i < maskedCount; i++ {
		builder.WriteRune('*')
	}

	for i := maskedCount; i < length; i++ {
		builder.WriteRune(runnes[i])
	}

	return builder.String(), nil
}

package utils

import (
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	var sb strings.Builder
	runnes := []rune(s)
	n := len(runnes)

	for i := 0; i < n; i++ {
		r := runnes[i]

		if unicode.IsUpper(r) {
			if i > 0 {
				prev := runnes[i-1]
				if unicode.IsLower(prev) || unicode.IsDigit(prev) {
					sb.WriteRune('_')
				} else if i+1 < n && unicode.IsLower(runnes[i+1]) {
					sb.WriteRune('_')
				}
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

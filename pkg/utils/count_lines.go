package utils

// CountLines counts lines separated by LF; empty text has zero lines and a trailing LF does not add a line.
func CountLines(s string) int {
	if s == "" {
		return 0
	}
	count := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			if i < len(s)-1 {
				count++
			}
		}
	}
	return count
}

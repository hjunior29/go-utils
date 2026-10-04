package utils

import "strings"

// WordCount counts groups of non-whitespace Unicode code points.
func WordCount(value string) int { return len(strings.Fields(value)) }

package utils

import "fmt"

// ClampInt limits value to inclusive bounds and rejects reversed bounds.
func ClampInt(value, lower, upper int) (int, error) {
	if lower > upper {
		return 0, fmt.Errorf("lower bound must not exceed upper bound")
	}
	if value < lower {
		return lower, nil
	}
	if value > upper {
		return upper, nil
	}
	return value, nil
}

package utils

// RotateString rotates Unicode code points left by a nonnegative offset modulo length.
func RotateString(s string, offset int) (string, error) {
	if offset < 0 {
		return "", ErrRotateStringNegativeOffset
	}

	runnes := []rune(s)
	if len(runnes) == 0 {
		return "", nil
	}

	shift := offset % len(runnes)
	if shift == 0 {
		return s, nil
	}

	rotated := append(runnes[shift:], runnes[:shift]...)
	return string(rotated), nil
}

var ErrRotateStringNegativeOffset = rotateStringErrorsNew("negative offset is not allowed")

func rotateStringErrorsNew(text string) error {
	return &rotateStringError{s: text}
}

type rotateStringError struct {
	s string
}

func (e *rotateStringError) Error() string {
	return e.s
}

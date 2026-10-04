package utils

import "testing"

func TestReverseString(t *testing.T) {
	for _, test := range []struct{ input, want string }{{"", ""}, {"a", "a"}, {"hello", "olleh"}, {"a😀é", "é😀a"}} {
		if got := ReverseString(test.input); got != test.want {
			t.Errorf("ReverseString(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

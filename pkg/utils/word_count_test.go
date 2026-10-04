package utils

import "testing"

func TestWordCount(t *testing.T) {
	for _, test := range []struct {
		input string
		want  int
	}{{"", 0}, {" \t\n", 0}, {"hello", 1}, {" hello  world\nagain ", 3}, {"one\u2003two", 2}} {
		if got := WordCount(test.input); got != test.want {
			t.Errorf("WordCount(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}

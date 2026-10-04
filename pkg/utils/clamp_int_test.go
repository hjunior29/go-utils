package utils

import "testing"

func TestClampInt(t *testing.T) {
	for _, test := range []struct{ value, lower, upper, want int }{{5, 0, 10, 5}, {-1, 0, 10, 0}, {11, 0, 10, 10}, {0, 0, 10, 0}, {10, 0, 10, 10}, {8, 3, 3, 3}} {
		got, err := ClampInt(test.value, test.lower, test.upper)
		if err != nil || got != test.want {
			t.Errorf("ClampInt(%d, %d, %d) = %d, %v; want %d", test.value, test.lower, test.upper, got, err, test.want)
		}
	}
	if _, err := ClampInt(5, 10, 0); err == nil {
		t.Fatal("expected reversed bounds to fail")
	}
}

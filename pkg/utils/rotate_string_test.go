package utils

import (
	"testing"
)

func TestRotateStringCases(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		offset    int
		want      string
		wantError bool
	}{
		{
			name:      "empty string",
			input:     "",
			offset:    3,
			want:      "",
			wantError: false,
		},
		{
			name:      "zero offset",
			input:     "hello",
			offset:    0,
			want:      "hello",
			wantError: false,
		},
		{
			name:      "normal rotation",
			input:     "abcdef",
			offset:    2,
			want:      "cdefab",
			wantError: false,
		},
		{
			name:      "offset greater than length",
			input:     "abc",
			offset:    5,
			want:      "cab",
			wantError: false,
		},
		{
			name:      "unicode characters",
			input:     "\u3053\u3093\u306b\u3061\u306f",
			offset:    2,
			want:      "\u306b\u3061\u306f\u3053\u3093",
			wantError: false,
		},
		{
			name:      "negative offset error",
			input:     "test",
			offset:    -1,
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RotateString(tt.input, tt.offset)
			if (err != nil) != tt.wantError {
				t.Errorf("RotateString() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("RotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

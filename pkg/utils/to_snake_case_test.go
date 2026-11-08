package utils

import (
	"testing"
)

func TestToSnakeCase_Table(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "EmptyString",
			input:    "",
			expected: "",
		},
		{
			name:     "SimpleCamelCase",
			input:    "camelCase",
			expected: "camel_case",
		},
		{
			name:     "SimplePascalCase",
			input:    "PascalCase",
			expected: "pascal_case",
		},
		{
			name:     "AcronymBoundary",
			input:    "JSONParser",
			expected: "json_parser",
		},
		{
			name:     "TrailingAcronym",
			input:    "GetHTTPResponse",
			expected: "get_http_response",
		},
		{
			name:     "AlreadySnakeCase",
			input:    "snake_case",
			expected: "snake_case",
		},
		{
			name:     "SingleLetter",
			input:    "A",
			expected: "a",
		},
		{
			name:     "DigitsInIdentifier",
			input:    "Version2Update",
			expected: "version2_update",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToSnakeCase(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

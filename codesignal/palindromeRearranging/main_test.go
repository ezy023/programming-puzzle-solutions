package main

import "testing"

func TestPalindromeRearranging(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "one",
			input:    "aabb",
			expected: true,
		},
		{
			name:     "three",
			input:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc",
			expected: false,
		},
		{
			name:     "four",
			input:    "abbcabb",
			expected: true,
		},
		{
			name:     "custom",
			input:    "baabb",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := palindromeRearranging(tc.input)
			if got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

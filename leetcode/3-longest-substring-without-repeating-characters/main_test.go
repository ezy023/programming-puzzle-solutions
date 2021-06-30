package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "one",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "two",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "three",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "four",
			input:    "",
			expected: 0,
		},
		{
			name:     "five",
			input:    " ",
			expected: 1,
		},
		{
			name:     "six",
			input:    "au",
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.input)
			if got != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}

package main

import "testing"

func TestFirstNotRepeatingCharacter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "one",
			input:    "abacabad",
			expected: "c",
		},
		{
			name:     "two",
			input:    "abacabaabacaba",
			expected: "_",
		},
		{
			name:     "three",
			input:    "z",
			expected: "z",
		},
		{
			name:     "four",
			input:    "bcb",
			expected: "c",
		},
		{
			name:     "six",
			input:    "abcdefghijklmnopqrstuvwxyziflskecznslkjfabe",
			expected: "d",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := firstNotRepeatingCharacter(tc.input)
			if tc.expected != got {
				t.Errorf("Expected %q, got %q", tc.expected, got)
			}
		})
	}
}

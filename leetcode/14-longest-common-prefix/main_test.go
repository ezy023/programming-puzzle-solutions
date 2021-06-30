package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name: "one",
			input: []string{
				"flower",
				"flow",
				"flight",
			},
			expected: "fl",
		},
		{
			name: "two",
			input: []string{
				"dog",
				"racecar",
				"car",
			},
			expected: "",
		},
		{
			name: "three",
			input: []string{
				"flower",
				"flow",
				"flowing",
				"flowering",
			},
			expected: "flow",
		},
		{
			name: "four",
			input: []string{
				"foo",
			},
			expected: "foo",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := longestCommonPrefix(tc.input)
			if tc.expected != got {
				t.Errorf("Expected %q, got %q", tc.expected, got)
			}
		})
	}
}

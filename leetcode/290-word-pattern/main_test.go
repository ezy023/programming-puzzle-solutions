package main

import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name       string
		pattern, s string
		expected   bool
	}{
		{
			name:     "one",
			pattern:  "abba",
			s:        "dog cat cat dog",
			expected: true,
		},
		{
			name:     "two",
			pattern:  "abba",
			s:        "dog cat cat fish",
			expected: false,
		},
		{
			name:     "three",
			pattern:  "aaaa",
			s:        "dog cat cat dog",
			expected: false,
		},
		{
			name:     "four",
			pattern:  "abba",
			s:        "dog dog dog dog",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := wordPattern(tc.pattern, tc.s)
			if tc.expected != got {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

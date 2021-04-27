package main

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "one",
			input:    "(1+(2*3)+((8)/4))+1",
			expected: 3,
		},
		{
			name:     "two",
			input:    "(1)+((2))+(((3)))",
			expected: 3,
		},
		{
			name:     "three",
			input:    "1+(2*3)/(2-1)",
			expected: 1,
		},
		{
			name:     "four",
			input:    "1",
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxDepth(tc.input)
			if tc.expected != got {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}

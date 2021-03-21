package main

import "testing"

func TestArrayMaximalAdjacentDifference(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "one",
			input:    []int{2, 4, 1, 0},
			expected: 3,
		},
		{
			name:     "two",
			input:    []int{1, 1, 1, 1},
			expected: 0,
		},
		{
			name:     "three",
			input:    []int{-1, 4, 10, 3, -2},
			expected: 7,
		},
		{
			name:     "four",
			input:    []int{10, 11, 13},
			expected: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := arrayMaximalAdjacentDifference(tc.input)
			if out != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, out)
			}
		})
	}
}

package main

import "testing"

func TestArrayChange(t *testing.T) {
	cases := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "zero",
			in:       []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "one",
			in:       []int{1, 1, 1},
			expected: 3,
		},
		{
			name:     "two",
			in:       []int{-1000, 0, -2, 0},
			expected: 5,
		},
		{
			name:     "three",
			in:       []int{2, 1, 10, 1},
			expected: 12,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := arrayChange(tc.in)
			if out != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, out)
			}
		})
	}
}

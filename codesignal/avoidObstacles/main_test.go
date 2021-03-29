package main

import "testing"

func TestAvoidObstacles(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "one",
			input:    []int{5, 3, 6, 7, 9},
			expected: 4,
		},
		{
			name:     "two",
			input:    []int{2, 3},
			expected: 4,
		},
		{
			name:     "three",
			input:    []int{1, 4, 10, 6, 2},
			expected: 7,
		},
		{
			name:     "four",
			input:    []int{1000, 999},
			expected: 6,
		},
		{
			name:     "five",
			input:    []int{19, 32, 11, 23},
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := avoidObstacles(tc.input)
			if got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

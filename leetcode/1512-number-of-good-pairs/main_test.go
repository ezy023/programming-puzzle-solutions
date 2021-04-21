package main

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "one",
			input:    []int{1, 2, 3, 1, 1, 3},
			expected: 4,
		},
		{
			name:     "two",
			input:    []int{1, 1, 1, 1},
			expected: 6,
		},
		{
			name:     "three",
			input:    []int{1, 2, 3},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := numIdenticalPairs(tc.input)
			if tc.expected != got {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}

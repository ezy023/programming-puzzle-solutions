package main

import "reflect"
import "testing"

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name            string
		input, expected []int
	}{
		{
			name:     "one",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "two",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "three",
			input:    []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			name:     "four",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "five",
			input:    []int{1, 3, 2},
			expected: []int{2, 1, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nextPermutation(tc.input)
			if !reflect.DeepEqual(tc.expected, tc.input) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.input)
			}
		})
	}
}

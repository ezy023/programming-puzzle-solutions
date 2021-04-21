package main

import "reflect"
import "testing"

func TestSmallerThanCurrent(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "one",
			input:    []int{8, 1, 2, 2, 3},
			expected: []int{4, 0, 1, 1, 3},
		},
		{
			name:     "two",
			input:    []int{6, 5, 4, 8},
			expected: []int{2, 1, 0, 3},
		},
		{
			name:     "three",
			input:    []int{7, 7, 7, 7},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := smallerNumbersThanCurrent(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

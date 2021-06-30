package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSums(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		input    []int
		expected []int
	}{
		{
			name:     "one",
			target:   9,
			input:    []int{2, 7, 11, 15},
			expected: []int{0, 1},
		},
		{
			name:     "two",
			target:   6,
			input:    []int{3, 2, 4},
			expected: []int{1, 2},
		},
		{
			name:     "three",
			target:   6,
			input:    []int{3, 3},
			expected: []int{0, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSums(tc.input, tc.target)
			sort.Ints(got)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

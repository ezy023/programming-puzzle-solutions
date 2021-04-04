package main

import "reflect"
import "testing"

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "one",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name:  "two",
			input: []int{0, 1},
			expected: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name:  "three",
			input: []int{1},
			expected: [][]int{
				{1},
			},
		},
		{
			name:  "four",
			input: []int{0, -1, 1},
			expected: [][]int{
				{-1, 0, 1},
				{-1, 1, 0},
				{0, -1, 1},
				{0, 1, -1},
				{1, -1, 0},
				{1, 0, -1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := permute(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected %v, got %v\n", tc.expected, got)
			}
		})
	}
}

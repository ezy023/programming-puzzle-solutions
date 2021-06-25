package main

import "testing"

func TestTreeDiameter(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		tree     [][]int
		expected int
	}{
		{
			name: "one",
			n:    10,
			tree: [][]int{
				{2, 5},
				{5, 7},
				{5, 1},
				{1, 9},
				{1, 0},
				{7, 6},
				{6, 3},
				{3, 8},
				{8, 4}},
			expected: 7,
		},
		{
			name:     "two",
			n:        1,
			tree:     [][]int{},
			expected: 0,
		},
		{
			name:     "three",
			n:        2,
			tree:     [][]int{{1, 0}},
			expected: 1,
		},
		{
			name: "four",
			n:    3,
			tree: [][]int{
				{1, 2},
				{2, 0}},
			expected: 2,
		},
		{
			name: "five",
			n:    5,
			tree: [][]int{
				{3, 0},
				{3, 4},
				{2, 3},
				{3, 1}},
			expected: 2,
		},
		{
			name: "seven",
			n:    10,
			tree: [][]int{
				{1, 3},
				{7, 3},
				{5, 3},
				{8, 7},
				{4, 1},
				{2, 3},
				{9, 4},
				{0, 8},
				{6, 8}},
			expected: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := treeDiameter(tc.n, tc.tree)
			if tc.expected != got {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

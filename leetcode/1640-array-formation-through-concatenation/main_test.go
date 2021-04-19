package main

import "testing"

func TestCanFormArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		pieces   [][]int
		expected bool
	}{
		{
			name:     "one",
			arr:      []int{85},
			pieces:   [][]int{{85}},
			expected: true,
		},
		{
			name:     "two",
			arr:      []int{15, 88},
			pieces:   [][]int{{88}, {15}},
			expected: true,
		},
		{
			name:     "three",
			arr:      []int{49, 18, 16},
			pieces:   [][]int{{16, 18, 49}},
			expected: false,
		},
		{
			name:     "four",
			arr:      []int{91, 4, 64, 78},
			pieces:   [][]int{{78}, {4, 64}, {91}},
			expected: true,
		},
		{
			name:     "five",
			arr:      []int{1, 3, 5, 7},
			pieces:   [][]int{{2, 4, 6, 8}},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := canFormArray(tc.arr, tc.pieces)
			if tc.expected != got {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

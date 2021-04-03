package main

import "reflect"
import "testing"

func TestShuffle(t *testing.T) {
	tests := []struct {
		name           string
		nums, expected []int
		n              int
	}{
		{
			name:     "one",
			nums:     []int{2, 5, 1, 3, 4, 7},
			n:        3,
			expected: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name:     "two",
			nums:     []int{1, 2, 3, 4, 4, 3, 2, 1},
			n:        4,
			expected: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name:     "three",
			nums:     []int{1, 1, 2, 2},
			n:        2,
			expected: []int{1, 2, 1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := shuffle(tc.nums, tc.n)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

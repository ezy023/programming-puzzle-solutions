package main

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "one",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "two",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			name:     "three",
			nums:     []int{1, -1, 0},
			k:        0,
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := subarraySum(tc.nums, tc.k)
			if got != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}

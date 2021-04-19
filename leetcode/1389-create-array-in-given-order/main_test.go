package main

import "testing"
import "reflect"

func TestCreateTargetArray(t *testing.T) {
	tests := []struct {
		name        string
		nums, index []int
		expected    []int
	}{
		{
			name:     "one",
			nums:     []int{0, 1, 2, 3, 4},
			index:    []int{0, 1, 2, 2, 1},
			expected: []int{0, 4, 1, 3, 2},
		},
		{
			name:     "two",
			nums:     []int{1, 2, 3, 4, 0},
			index:    []int{0, 1, 2, 3, 0},
			expected: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := createTargetArray(tc.nums, tc.index)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("Expected %#v, got %#v", tc.expected, got)
			}
		})
	}
}

package main

import "testing"

func TestRestoreString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		indices  []int
		expected string
	}{
		{
			name:     "one",
			s:        "codeleet",
			indices:  []int{4, 5, 6, 7, 0, 2, 1, 3},
			expected: "leetcode",
		},
		{
			name:     "two",
			s:        "abc",
			indices:  []int{0, 1, 2},
			expected: "abc",
		},
		{
			name:     "three",
			s:        "aiohn",
			indices:  []int{3, 1, 4, 2, 0},
			expected: "nihao",
		},
		{
			name:     "four",
			s:        "aaiougrt",
			indices:  []int{4, 0, 2, 6, 7, 3, 1, 5},
			expected: "arigatou",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := restoreString(tc.s, tc.indices)
			if tc.expected != got {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

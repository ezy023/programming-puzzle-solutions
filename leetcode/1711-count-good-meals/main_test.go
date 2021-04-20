package main

import "testing"

func TestCountPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "one",
			input:    []int{1, 3, 5, 7, 9},
			expected: 4,
		},
		{
			name:     "two",
			input:    []int{1, 1, 1, 3, 3, 3, 7},
			expected: 15,
		},
		{
			name:     "three",
			input:    []int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234},
			expected: 12,
		},
		{
			name:     "four",
			input:    []int{1048576, 1048576},
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := countPairs(tc.input)
			if tc.expected != got {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}

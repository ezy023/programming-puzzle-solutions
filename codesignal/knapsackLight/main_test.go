package main

import "testing"

func TestKnapsackLight(t *testing.T) {
	tests := []struct {
		name                 string
		v1, w1, v2, w2, maxW int
		expected             int
	}{
		{
			name:     "one",
			v1:       10,
			w1:       5,
			v2:       6,
			w2:       4,
			maxW:     8,
			expected: 10,
		},
		{
			name:     "two",
			v1:       10,
			w1:       5,
			v2:       6,
			w2:       4,
			maxW:     9,
			expected: 16,
		},
		{
			name:     "three",
			v1:       5,
			w1:       3,
			v2:       7,
			w2:       4,
			maxW:     6,
			expected: 7,
		},
		{
			name:     "five",
			v1:       15,
			w1:       2,
			v2:       20,
			w2:       3,
			maxW:     2,
			expected: 15,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := knapsackLight(tc.v1, tc.w1, tc.v2, tc.w2, tc.maxW)
			if got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

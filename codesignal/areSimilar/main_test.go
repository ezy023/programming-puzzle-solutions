package main

import "testing"

func TestAreSimilar(t *testing.T) {
	cases := []struct {
		name     string
		a, b     []int
		expected bool
	}{
		{
			name:     "one",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "one",
			a:        []int{1, 2, 3},
			b:        []int{2, 1, 3},
			expected: true,
		},
		{
			name:     "one",
			a:        []int{1, 2, 2},
			b:        []int{2, 1, 1},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := areSimilar(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

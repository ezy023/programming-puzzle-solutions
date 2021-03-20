package main

import (
	"testing"
)

func TestAreAlmostEqual(t *testing.T) {
	cases := []struct {
		s1, s2   string
		expected bool
	}{
		{
			s1:       "bank",
			s2:       "kanb",
			expected: true,
		},
		{
			s1:       "attack",
			s2:       "defend",
			expected: false,
		},
		{
			s1:       "kelb",
			s2:       "kelb",
			expected: true,
		},
		{
			s1:       "abcd",
			s2:       "dcba",
			expected: false,
		},
		{
			s1:       "aa",
			s2:       "ac",
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			actual := areAlmostEqual(tc.s1, tc.s2)

			if actual != tc.expected {
				t.Errorf("Expected %v, got %v for s1: %s s2: %s", tc.expected, actual, tc.s1, tc.s2)
			}
		})
	}
}

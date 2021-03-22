package main

import "testing"

func TestLateRide(t *testing.T) {
	cases := []struct {
		name            string
		input, expected int
	}{
		{
			name:     "one",
			input:    240,
			expected: 4,
		},
		{
			name:     "two",
			input:    808,
			expected: 14,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := lateRide(tc.input)
			if out != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, out)
			}
		})
	}
}

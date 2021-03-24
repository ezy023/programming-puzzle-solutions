package main

import "testing"

func TestPhoneCall(t *testing.T) {
	tests := []struct {
		name                    string
		min1, min2_10, min11, s int
		expected                int
	}{
		{
			name:     "one",
			min1:     3,
			min2_10:  1,
			min11:    2,
			s:        20,
			expected: 14,
		},
		{
			name:     "two",
			min1:     2,
			min2_10:  2,
			min11:    1,
			s:        2,
			expected: 1,
		},
		{
			name:     "three",
			min1:     10,
			min2_10:  1,
			min11:    2,
			s:        22,
			expected: 11,
		},
		{
			name:     "four",
			min1:     2,
			min2_10:  2,
			min11:    1,
			s:        24,
			expected: 14,
		},
		{
			name:     "five",
			min1:     1,
			min2_10:  2,
			min11:    1,
			s:        6,
			expected: 3,
		},
		{
			name:     "six",
			min1:     10,
			min2_10:  10,
			min11:    10,
			s:        8,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := phoneCall(tc.min1, tc.min2_10, tc.min11, tc.s)
			if out != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, out)
			}
		})
	}
}

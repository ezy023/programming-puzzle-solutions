package main

import "testing"

func TestQueryString(t *testing.T) {
	tests := []struct {
		name     string
		S        string
		N        int
		expected bool
	}{
		{
			name:     "one",
			S:        "0110",
			N:        3,
			expected: true,
		},
		{
			name:     "two",
			S:        "0110",
			N:        4,
			expected: false,
		},
		{
			name:     "three",
			S:        "110101011011000011011111000000",
			N:        15,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := queryString(tc.S, tc.N)
			if tc.expected != got {
				t.Errorf("[queryString] Expected %v, got %v", tc.expected, got)
			}

			gotStrFmt := queryStringStringFormatting(tc.S, tc.N)
			if tc.expected != gotStrFmt {
				t.Errorf("[queryStringStringFormatting] Expected %v, got %v", tc.expected, gotStrFmt)
			}
		})
	}
}

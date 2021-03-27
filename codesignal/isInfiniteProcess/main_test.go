package main

import "testing"

func TestIsInfiniteProcess(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		exp  bool
	}{
		{
			name: "one",
			a:    2,
			b:    6,
			exp:  false,
		},
		{
			name: "one",
			a:    2,
			b:    3,
			exp:  true,
		},
		{
			name: "one",
			a:    2,
			b:    10,
			exp:  false,
		},
		{
			name: "four",
			a:    0,
			b:    1,
			exp:  true,
		},
		{
			name: "seven",
			a:    5,
			b:    10,
			exp:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isInfiniteProcess(tc.a, tc.b)
			if got != tc.exp {
				t.Errorf("Expected %v, got %v\n", tc.exp, got)
			}
		})
	}
}

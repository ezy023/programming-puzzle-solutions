package main

import "testing"

func TestAlmsotIncreasingSequence(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		out  bool
	}{
		{
			name: "one",
			in:   []int{1, 3, 2, 1},
			out:  false,
		},
		{
			name: "two",
			in:   []int{1, 3, 2},
			out:  true,
		},
		{
			name: "three",
			in:   []int{1, 2, 1, 2},
			out:  false,
		},
		{
			name: "four",
			in:   []int{3, 6, 5, 8, 10, 20, 15},
			out:  false,
		},
		{
			name: "five",
			in:   []int{1, 1, 2, 3, 4, 4},
			out:  false,
		},
		{
			name: "six",
			in:   []int{10, 1, 2, 3, 4, 5},
			out:  true,
		},
		{
			name: "seven",
			in:   []int{1, 2, 5, 3, 5},
			out:  true,
		},
		{
			name: "eight",
			in:   []int{1, 2, 3, 4, 3, 6},
			out:  true,
		},
	}

	for _, tc := range cases {
		result := almostIncreasingSequence(tc.in)
		t.Run(tc.name, func(t *testing.T) {
			if result != tc.out {
				t.Errorf("Expected %v, got %v for in: %#v", tc.out, result, tc.in)
			}
		})
	}
}

package main

import "testing"

func TestMatrixElementsSum(t *testing.T) {
	cases := []struct {
		name   string
		matrix [][]int
		sum    int
	}{
		{
			name: "one",
			matrix: [][]int{
				{0, 1, 1, 2},
				{0, 5, 0, 0},
				{2, 0, 3, 3}},
			sum: 9,
		},
		{
			name: "two",
			matrix: [][]int{
				{1, 1, 1, 0},
				{0, 5, 0, 1},
				{2, 1, 3, 10}},
			sum: 9,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := matrixElementsSum(tc.matrix)
			if out != tc.sum {
				t.Errorf("Expected %d, got %d", tc.sum, out)
			}
		})
	}
}

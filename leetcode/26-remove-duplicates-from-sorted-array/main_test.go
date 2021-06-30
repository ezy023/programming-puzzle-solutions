package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		expectedNums []int
		expectedK    int
	}{
		{
			name:         "one",
			input:        []int{1, 1, 2},
			expectedNums: []int{1, 2},
			expectedK:    2,
		},
		{
			name:         "one",
			input:        []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
			expectedK:    5,
		},
	}

	compareSlice := func(exp, got []int) bool {
		for idx, val := range exp {
			if got[idx] != val {
				return false
			}
		}
		return true
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := removeDuplicates(tc.input)
			if tc.expectedK != got {
				t.Errorf("Expected k %d, got %d", tc.expectedK, got)
			}

			if !compareSlice(tc.expectedNums, tc.input) {
				t.Errorf("Expected %v, got %v", tc.expectedNums, tc.input)
			}
		})
	}
}

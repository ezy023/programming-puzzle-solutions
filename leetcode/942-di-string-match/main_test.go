package main

import "testing"
import "reflect"

func TestDIStringMatch(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "one",
			input:    "IDID",
			expected: []int{0, 4, 1, 3, 2},
		},
		{
			name:     "two",
			input:    "III",
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "three",
			input:    "DDI",
			expected: []int{3, 2, 0, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := diStringMatch(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

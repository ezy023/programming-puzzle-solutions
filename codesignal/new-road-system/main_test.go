package main

import "testing"

func TestNewRoadSystem(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]bool
		expected bool
	}{
		{
			name: "one",
			input: [][]bool{
				{false, true, false, false},
				{false, false, true, false},
				{true, false, false, true},
				{false, false, true, false},
			},
			expected: true,
		},
		{
			name: "two",
			input: [][]bool{
				{false, true, false, false, false, false, false},
				{true, false, false, false, false, false, false},
				{false, false, false, true, false, false, false},
				{false, false, true, false, false, false, false},
				{false, false, false, false, false, false, true},
				{false, false, false, false, true, false, false},
				{false, false, false, false, false, true, false},
			},
			expected: true,
		},
		{
			name: "three",
			input: [][]bool{
				{false, true, false},
				{false, false, false},
				{true, false, false},
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := newRoadSystem(tc.input)
			if got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

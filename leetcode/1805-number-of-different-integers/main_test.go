package main

import "testing"

func TestNumDifferentIntegers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "one",
			input:    "a123bc34d8ef34",
			expected: 3,
		},
		{
			name:     "two",
			input:    "leet1234code234",
			expected: 2,
		},
		{
			name:     "three",
			input:    "a1b01c001",
			expected: 1,
		},
		{
			name:     "four",
			input:    "0a0",
			expected: 1,
		},
		{
			name:     "five",
			input:    "uu717761265362523668772526716127260222200144937319826j717761265362523668772526716127260222200144937319826k717761265362523668772526716127260222200144937319826b7177612653625236687725267161272602222001449373198262hgb9utohfvfrxprqva3y5cdfdudf7zuh64mobfr6mhy17",
			expected: 9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := numDifferentIntegers(tc.input)
			if got != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}

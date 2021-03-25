package main

import "testing"

func TestReachNextLevel(t *testing.T) {
	tests := []struct {
		name                          string
		experience, threshold, reward int
		expected                      bool
	}{
		{
			name:       "one",
			experience: 10,
			threshold:  15,
			reward:     5,
			expected:   true,
		},
		{
			name:       "two",
			experience: 10,
			threshold:  15,
			reward:     5,
			expected:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := reachNextLevel(tc.experience, tc.threshold, tc.reward)
			if got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

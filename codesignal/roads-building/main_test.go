package main

import "testing"
import "reflect"

func TestRoadsBuilding(t *testing.T) {
	tests := []struct {
		name            string
		cities          int
		roads, expected [][]int
	}{
		{
			name:     "one",
			cities:   4,
			roads:    [][]int{{0, 1}, {1, 2}, {2, 0}},
			expected: [][]int{{0, 3}, {1, 3}, {2, 3}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := roadsBuilding(tc.cities, tc.roads)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

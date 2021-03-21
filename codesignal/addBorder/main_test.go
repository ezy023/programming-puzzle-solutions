package main

import "testing"
import "reflect"

func TestAddBorder(t *testing.T) {
	cases := []struct {
		name     string
		picture  []string
		expected []string
	}{
		{
			name: "one",
			picture: []string{
				"abc",
				"ded",
			},
			expected: []string{
				"*****",
				"*abc*",
				"*ded*",
				"*****",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := addBorder(tc.picture)
			if !reflect.DeepEqual(out, tc.expected) {
				t.Errorf("Expected %#v\n, go %#v", tc.expected, out)
			}
		})
	}
}

package main

import "testing"
import "reflect"

func TestAccountsMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		{
			name: "one",
			input: [][]string{
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
		},
		{
			name: "two",
			input: [][]string{
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
				{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
				{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
				{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
				{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
			},
			expected: [][]string{
				{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
				{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
				{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
				{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
			},
		},
		{
			name: "three",
			input: [][]string{
				{"David", "David0@m.co", "David1@m.co"},
				{"David", "David3@m.co", "David4@m.co"},
				{"David", "David4@m.co", "David5@m.co"},
				{"David", "David2@m.co", "David3@m.co"},
				{"David", "David1@m.co", "David2@m.co"},
			},
			expected: [][]string{
				{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotAccounts := accountsMerge(tc.input)
			var fail bool = true
			for _, account := range gotAccounts {
				name := account[0]
				var expAccount []string
				for _, a := range tc.expected {
					if a[0] == name {
						expAccount = a
					}
				}
				if !reflect.DeepEqual(expAccount, account) {
					fail = true
				} else {
					fail = false
					continue
				}
			}
			if fail {
				t.Errorf("Expected %v\n\t\tGot %v", tc.expected, gotAccounts)
			}
		})
	}
}

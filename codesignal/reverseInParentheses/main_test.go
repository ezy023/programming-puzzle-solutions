package main

import "testing"

func TestReverseInParentheses(t *testing.T) {
	cases := []struct {
		name, input, expected string
	}{
		{
			name:     "one",
			input:    "(bar)",
			expected: "rab",
		},
		{
			name:     "two",
			input:    "foo(bar)baz",
			expected: "foorabbaz",
		},
		{
			name:     "three",
			input:    "foo(bar)baz(blim)",
			expected: "foorabbazmilb",
		},
		{
			name:     "four",
			input:    "foo(bar(baz))blim",
			expected: "foobazrabblim",
		},
		{
			name:     "six",
			input:    "()",
			expected: "",
		},
		{
			name:     "seven",
			input:    "(abc)d(efg)",
			expected: "cbadgfe",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := reverseInParentheses(tc.input)
			if out != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, out)
			}
		})
	}
}

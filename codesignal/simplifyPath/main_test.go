package main

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name           string
		path, expected string
	}{
		{
			name:     "one",
			path:     "/home/a/./x/../b//c/",
			expected: "/home/a/b/c",
		},
		{
			name:     "two",
			path:     "/a/b/c/../..",
			expected: "/a",
		},
		{
			name:     "three",
			path:     "/../",
			expected: "/",
		},
		{
			name:     "four",
			path:     "/",
			expected: "/",
		},
		{
			name:     "five",
			path:     "//a//b//./././c",
			expected: "/a/b/c",
		},
		{
			name:     "five",
			path:     "a/../../b/",
			expected: "/b",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := simplifyPath(tc.path)
			if got != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, got)
			}
		})
	}
}

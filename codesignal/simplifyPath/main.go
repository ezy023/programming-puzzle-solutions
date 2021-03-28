package main

import "strings"

func simplifyPath(path string) string {
	parts := make([]string, 0)
	for _, part := range strings.Split(path, "/") {
		if part == ".." {
			if len(parts) > 0 {
				parts = parts[:len(parts)-1]
			}
		} else if part == "" || part == "." {
			continue
		} else {
			parts = append(parts, part)
		}
	}
	if len(parts) == 0 {
		parts = []string{"/"}
	} else if parts[0] != "/" {
		parts = append([]string{""}, parts...)
	}
	return strings.Join(parts, "/")
}

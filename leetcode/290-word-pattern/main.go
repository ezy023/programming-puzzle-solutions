package main

import "strings"

func wordPattern(pattern, s string) bool {
	chars := make(map[byte]string)
	words := make(map[string]bool)
	parts := strings.Split(s, " ")
	if len(pattern) != len(parts) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if word, ok := chars[pattern[i]]; !ok {
			if words[parts[i]] {
				return false
			}
			chars[pattern[i]] = parts[i]
			words[parts[i]] = true // seen
		} else {
			if word != parts[i] {
				return false
			}
		}
	}

	return true
}

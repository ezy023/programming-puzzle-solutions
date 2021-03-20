package main

import (
	_ "fmt"
)

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	if len(s1) != len(s2) {
		return false
	}

	diffIdx := []int{}
	for idx, c := range s1 {
		if c != rune(s2[idx]) {
			diffIdx = append(diffIdx, idx)
		}
	}

	return len(diffIdx) == 2 && s1[diffIdx[0]] == s2[diffIdx[1]] && s1[diffIdx[1]] == s2[diffIdx[0]]
}

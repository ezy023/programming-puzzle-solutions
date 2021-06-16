package main

import "fmt"

func firstNotRepeatingCharacter(s string) string {
	deleteFrom := func(col []rune, c rune) []rune {
		for idx, item := range col {
			if item == c {
				return append(col[:idx], col[idx+1:]...)
			}
		}
		return col
	}

	seen := make(map[rune]bool)
	stack := make([]rune, 0)
	for _, r := range s {
		if _, ok := seen[r]; ok {
			stack = deleteFrom(stack, r)
		} else {
			stack = append(stack, r)
		}

		seen[r] = true
	}

	if len(stack) == 0 {
		return "_"
	} else {
		return string(stack[0])
	}
}

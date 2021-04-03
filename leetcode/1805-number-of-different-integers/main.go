package main

import "strconv"

func numDifferentIntegers(word string) int {
	seen := make(map[float64]bool)
	buf := make([]rune, 0, 0)
	for _, r := range word {
		if int(r-'0') >= 0 && int(r-'0') <= 9 {
			buf = append(buf, r)
		} else {
			if len(buf) > 0 {
				i, _ := strconv.ParseFloat(string(buf), 64)
				seen[i] = true
				buf = []rune{}
			}
		}
	}
	if len(buf) > 0 {
		i, _ := strconv.ParseFloat(string(buf), 64)
		seen[i] = true
	}

	return len(seen)
}

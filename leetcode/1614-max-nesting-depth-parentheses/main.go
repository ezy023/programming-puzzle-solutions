package main

func maxDepth(s string) int {
	max := 0
	depth := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			depth++
			if depth > max {
				max = depth
			}
		} else if s[i] == ')' {
			depth--
		}
	}

	return max
}

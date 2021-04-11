package main

import "fmt"
import "strings"

func queryStringStringFormatting(S string, N int) bool {
	base2 := func(n int) string {
		return fmt.Sprintf("%b", n)
	}

	for i := 1; i <= N; i++ {
		if !strings.Contains(S, base2(i)) {
			return false
		}
	}
	return true
}

func queryString(S string, N int) bool {
	base2 := func(n int) string {
		if n == 0 {
			return "0"
		}
		stack := make([]int, 0, 0)
		for n > 0 {
			i := n % 2
			stack = append(stack, i)
			n = n / 2
		}
		s := ""
		for i := len(stack) - 1; i >= 0; i-- {
			s += fmt.Sprintf("%d", stack[i])
		}
		return s
	}

	for i := 1; i <= N; i++ {
		res := strings.Contains(S, base2(i))
		if !res {
			return false
		}
	}
	return true
}

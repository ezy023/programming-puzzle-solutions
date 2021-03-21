package main

func reverseInParentheses(inputString string) string {
	reverse := func(in string) string {
		b := make([]byte, 0, len(in))
		for i := len(in) - 1; i >= 0; i-- {
			b = append(b, in[i])
		}
		return string(b)
	}

	collector := make([]byte, 0, len(inputString))

	collectInParens := func(in string, start int) int {
		count := 1
		for i := start; i < len(in); i++ {
			if in[i] == '(' {
				count++
			}
			if in[i] == ')' {
				count--
			}
			if count == 0 {
				return i
			}
		}
		return -1 //unbalanced parens
	}

	for i := 0; i < len(inputString); {
		if inputString[i] == '(' {
			endIdx := collectInParens(inputString, i+1)
			substr := inputString[i+1 : endIdx]
			reversed := reverse(reverseInParentheses(substr))
			for j := 0; j < len(reversed); j++ {
				collector = append(collector, reversed[j])
			}
			i = endIdx + 1
		} else {
			collector = append(collector, inputString[i])
			i++
		}

	}

	return string(collector)
}

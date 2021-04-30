package main

func diStringMatch(s string) []int {
	start, end := 0, len(s)
	buf := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			buf[i] = start
			start++
		} else if s[i] == 'D' {
			buf[i] = end
			end--
		}
	}

	buf[len(s)] = start
	return buf
}

package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	buf := make([]byte, 0)
	first := strs[0]
	for c := 0; c < len(first); c++ {
		for i := 1; i < len(strs); i++ {
			if c >= len(strs[i]) {
				return string(buf)
			}
			if strs[i][c] != first[c] {
				return string(buf)
			}
		}
		buf = append(buf, first[c])
	}

	return string(buf)
}

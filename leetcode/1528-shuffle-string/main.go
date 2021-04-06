package main

func restoreString(s string, indices []int) string {
	buf := make([]byte, len(s))
	for i, idx := range indices {
		buf[idx] = s[i]
	}
	return string(buf)
}

package main

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	max := 0
	var start int
	var end int
	for start = 0; start < len(s)-1; start++ {
		seen := make(map[byte]bool)
		seen[s[start]] = true
		for end = start + 1; end < len(s); end++ {
			if ok := seen[s[end]]; ok {
				break
			}
			seen[s[end]] = true
		}
		length := end - start
		if length > max {
			max = length
		}
	}

	return max
}

package main

func palindromeRearranging(inputString string) bool {
	counter := make(map[rune]int)
	for _, c := range inputString {
		counter[c]++
	}

	singleCharAllowed := len(inputString) % 2
	for _, v := range counter {
		if v%2 == 1 {
			if singleCharAllowed < 1 {
				return false
			} else {
				singleCharAllowed--
			}
		} else if v%2 != 0 {
			return false
		}
	}
	return true
}

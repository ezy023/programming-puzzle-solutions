package main

func almostIncreasingSequence(sequence []int) bool {
	skipped := 0
	for i := 1; i < len(sequence); i++ {
		if sequence[i-1] >= sequence[i] {
			skipped++
			if skipped > 1 {
				return false
			}
			if i >= 2 && i+1 < len(sequence) && sequence[i-2] >= sequence[i] && sequence[i-1] >= sequence[i+1] {
				return false
			}
		}
	}
	return true
}

package main

func areSimilar(a, b []int) bool {
	diffIdx := make([]int, 0, 2)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffIdx = append(diffIdx, i)
		}
	}

	if len(diffIdx) == 0 {
		return true
	}
	if len(diffIdx) == 2 {
		if a[diffIdx[0]] == b[diffIdx[1]] && a[diffIdx[1]] == b[diffIdx[0]] {
			return true
		}
	}

	return false
}

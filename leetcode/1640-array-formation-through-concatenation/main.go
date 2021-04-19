package main

func canFormArray(arr []int, pieces [][]int) bool {
	var a int = 0

	for i := 0; i < len(pieces); {
		if arr[a] == pieces[i][0] {
			a++
			for j := 1; j < len(pieces[i]); j++ {
				if arr[a] != pieces[i][j] {
					return false
				}
				a++
			}
			i = 0
		} else {
			i++
		}
		if a == len(arr) {
			return true
		}
	}

	return false
}

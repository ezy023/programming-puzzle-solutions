package main

func matrixElementsSum(matrix [][]int) int {
	contains := func(col []int, n int) bool {
		for _, i := range col {
			if i == n {
				return true
			}
		}
		return false
	}

	badCol := make([]int, 0, 0)
	sum := 0
	for _, floor := range matrix {
		for idx, room := range floor {
			if contains(badCol, idx) {
				continue
			}
			if room == 0 {
				badCol = append(badCol, idx)
			}
			sum += room
		}
	}
	return sum
}

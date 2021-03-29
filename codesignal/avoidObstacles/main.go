package main

func avoidObstacles(inputArray []int) int {
	allmod := func(in []int, d int) bool {
		for _, v := range in {
			if v%d == 0 {
				return false
			}
		}
		return true
	}

	var i int
	for i = 2; i < 1000; i++ {
		if allmod(inputArray, i) {
			return i
		}
	}

	return i + 1
}

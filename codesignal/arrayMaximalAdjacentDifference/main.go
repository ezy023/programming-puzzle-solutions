package main

import "math"

func arrayMaximalAdjacentDifference(inputArray []int) int {
	max := 0
	for i := 1; i < len(inputArray); i++ {
		diff := int(math.Abs(float64(inputArray[i] - inputArray[i-1])))
		if diff > max {
			max = diff
		}
	}

	return max
}

package main

func arrayChange(inputArray []int) int {
	count := 0
	for i := 1; i < len(inputArray); i++ {
		if inputArray[i-1] >= inputArray[i] {
			diff := inputArray[i-1] - inputArray[i] + 1
			inputArray[i] = inputArray[i-1] + 1
			count += diff
		}
	}
	return count
}

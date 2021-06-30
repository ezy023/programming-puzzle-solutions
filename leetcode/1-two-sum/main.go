package main

func twoSums(nums []int, target int) []int {
	out := make([]int, 2)
outer:
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				out[0], out[1] = i, j
				break outer
			}
		}
	}
	return out
}

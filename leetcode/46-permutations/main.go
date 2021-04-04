package main

import "sort"

func permute(nums []int) [][]int {
	out := make([][]int, 0, 0)
	sort.Ints(nums)
	buf := make([]int, len(nums))
	copy(buf, nums)
	out = append(out, buf)
	for {
		var k int = -1
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				k = i
			}
		}
		if k == -1 {
			break
		}

		var l int
		for i := k + 1; i < len(nums); i++ {
			if nums[k] < nums[i] {
				l = i
			}
		}
		nums[k], nums[l] = nums[l], nums[k]
		for i, j := k+1, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		buf := make([]int, len(nums))
		copy(buf, nums)
		out = append(out, buf)
	}
	return out
}

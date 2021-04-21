package main

import "sort"

func smallerNumbersThanCurrentOne(nums []int) []int {
	buf := make([]int, len(nums))
	copy(buf, nums)
	sort.Ints(buf)
	out := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		current := nums[i]
		for j := 0; j < len(buf); j++ {
			if buf[j] < current {
				out[i]++
			}
		}
	}
	return out
}

func smallerNumbersThanCurrent(nums []int) []int {
	out := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				out[i]++
			}
		}
	}

	return out
}

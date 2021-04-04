package main

// Algorithm for solution explained: https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func nextPermutation(nums []int) {
	var k int = -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			k = i
		}
	}

	var l int
	for i := k + 1; i < len(nums) && k != -1; i++ {
		if nums[k] < nums[i] {
			l = i
		}
	}
	if k != -1 {
		nums[k], nums[l] = nums[l], nums[k]
	}
	for i, j := k+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

package main

func subarraySum(nums []int, k int) int {
	var totals int = 0
	for i := 0; i < len(nums); i++ {
		var sum int = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				totals++
			}
		}
	}
	return totals
}

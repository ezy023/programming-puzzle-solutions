package main

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	tmp := nums[0]
	insertIdx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != tmp {
			nums[insertIdx] = nums[i]
			insertIdx++
			tmp = nums[i]
		}
	}

	return insertIdx
}

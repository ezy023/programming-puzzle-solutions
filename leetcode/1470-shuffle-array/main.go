package main

func shuffle(nums []int, n int) []int {
	buf := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		buf = append(buf, nums[i])
		buf = append(buf, nums[n+i])
	}
	return buf
}

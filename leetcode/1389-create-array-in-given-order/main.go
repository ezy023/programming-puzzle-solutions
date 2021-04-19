package main

func createTargetArray(nums []int, index []int) []int {
	insert := func(arr []int, idx, item int) {
		if arr[idx] < 0 {
			arr[idx] = item
		} else {
			for i := len(arr) - 2; i >= idx; i-- {
				arr[i+1] = arr[i]
			}
			arr[idx] = item
		}
	}

	out := make([]int, len(nums))
	for i := 0; i < len(index); i++ {
		insert(out, index[i], nums[i])
	}

	return out
}

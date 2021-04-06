package main

import "fmt"
import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	arr := make([]int, 0, 0)
	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}
	fmt.Printf("arr = %#v\n", arr)
	var num float64 = 0
	var pos float64 = 0
	for i := len(arr) - 1; i >= 0; i, pos = i-1, pos+1 {
		num += math.Pow(2, pos) * float64(arr[i])
	}

	return int(num)
}

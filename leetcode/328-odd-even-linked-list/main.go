package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	odd := head
	var even *ListNode = head.Next
	var evenHead *ListNode = even

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%#v ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("")
}

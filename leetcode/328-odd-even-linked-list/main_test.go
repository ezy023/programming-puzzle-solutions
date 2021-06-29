package main

import "testing"
import "reflect"

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "one",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 5, 2, 4},
		},
		{
			name:     "two",
			input:    []int{2, 1, 3, 5, 6, 4, 7},
			expected: []int{2, 3, 6, 7, 1, 5, 4},
		},
		{
			name:     "three",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inputLinked := toLinkedList(tc.input)
			printList(inputLinked)
			got := oddEvenList(inputLinked)
			gotList := fromLinkedList(got)

			if !reflect.DeepEqual(tc.expected, gotList) {
				t.Errorf("Expected %v, got %v", tc.expected, gotList)
			}
		})
	}
}

func toLinkedList(in []int) *ListNode {
	if len(in) < 1 {
		return nil
	}
	head := &ListNode{
		Val:  in[0],
		Next: nil,
	}
	cur := head
	for _, v := range in[1:] {
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		cur.Next = node
		cur = cur.Next
	}

	return head
}

func fromLinkedList(head *ListNode) []int {
	buf := make([]int, 0)
	cur := head
	for cur != nil {
		buf = append(buf, cur.Val)
		cur = cur.Next
	}

	return buf
}

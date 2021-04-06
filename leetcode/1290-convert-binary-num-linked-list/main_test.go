package main

import "fmt"
import "testing"

func TestGetDecimalValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "one",
			input:    []int{1, 0, 1},
			expected: 5,
		},
		{
			name:     "two",
			input:    []int{0},
			expected: 0,
		},
		{
			name:     "three",
			input:    []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			expected: 18880,
		},
		{
			name:     "four",
			input:    []int{0, 0},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var head *ListNode = toLinkedList(tc.input)
			// printLinkedList(head)
			got := getDecimalValue(head)
			if tc.expected != got {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func toLinkedList(in []int) *ListNode {
	head := &ListNode{Val: in[0], Next: nil}
	current := head
	for i := 1; i < len(in); i++ {
		current.Next = &ListNode{Val: in[i], Next: nil}
		current = current.Next
	}
	return head
}

func printLinkedList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("Value %d\n", p.Val)
		p = p.Next
	}
}

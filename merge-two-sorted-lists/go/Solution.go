package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head, curr, left, right *ListNode
	left = l1
	right = l2
	if left.Val < right.Val {
		head, left = left, left.Next
	} else {
		head, right = right, right.Next
	}
	curr = head
	for left != nil && right != nil {
		if left.Val < right.Val {
			curr.Next, left = left, left.Next
		} else {
			curr.Next, right = right, right.Next
		}
		curr = curr.Next
	}
	if left != nil {
		curr.Next = left
	} else if right != nil {
		curr.Next = right
	}

	return head
}

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	curr := mergeTwoLists(l1, l2)
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

package main

import (
	"fmt"
)

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sum(l1 *ListNode, l2 *ListNode, overflow int) *ListNode {
	if l1 == nil && l2 == nil {
		if overflow > 0 {
			return &ListNode{Val: overflow}
		} else {
			return nil
		}
	}
	if l1 == nil {
		l1 = &ListNode{Val: 0}
	} else if l2 == nil {
		l2 = &ListNode{Val: 0}
	}
	// assert(l1.Val < 10)
	// assert(l2.Val < 10)
	s := &ListNode{Val: l1.Val + l2.Val + overflow}
	if s.Val < 10 {
		overflow = 0
	} else {
		overflow = 1
		s.Val %= 10
	}
	s.Next = sum(l1.Next, l2.Next, overflow)
	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s := sum(l1, l2, 0)
	return s
}

func main() {
	left := &ListNode{Val: 2}
	left.Next = &ListNode{Val: 4}
	left.Next.Next = &ListNode{Val: 3}

	right := &ListNode{Val: 5}
	right.Next = &ListNode{Val: 6}
	right.Next.Next = &ListNode{Val: 4}

	sum := addTwoNumbers(left, right)
	for sum != nil {
		fmt.Printf(" %d", sum.Val)
		sum = sum.Next
	}
	fmt.Println()
}

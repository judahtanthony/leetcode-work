package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	node := head
	var reverse *ListNode
	for node != nil {
		reverse = &ListNode{node.Val, reverse}
		node = node.Next
	}

	return reverse
}

func main() {
	reversed := reverseList(&ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	})
	for n := reversed; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
}

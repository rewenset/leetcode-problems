package main

import (
	"fmt"
)

//ListNode represents element of singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func linkedListAsInts(node *ListNode) []int {
	var s []int

	for e := node; e != nil; e = e.Next {
		s = append(s, e.Val)
	}

	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{}
	curr := first

	var a, r int
	for {
		if l1 != nil {
			r = l1.Val
			l1 = l1.Next
		} else {
			r = 0
		}

		if l2 != nil {
			r += l2.Val
			l2 = l2.Next
		}
		r += a

		a = 0

		if r >= 10 {
			a = 1
			r -= 10
		}
		curr.Val = r

		if l1 == nil && l2 == nil {
			break
		}

		curr.Next = &ListNode{}
		curr = curr.Next
	}
	if a != 0 {
		curr.Next = &ListNode{Val: 1}
	}
	return first
}

func main() {
	fmt.Println(linkedListAsInts(addTwoNumbers(
		&ListNode{9, &ListNode{9, nil}},
		// &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		&ListNode{1, nil},
		// &ListNode{5, nil},
	)))
}

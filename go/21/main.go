package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func print(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Print("\n")
}

func create(nums []int) *ListNode {
	head := &ListNode{}

	current := head
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return head.Next
}

func main() {
	result := mergeTwoLists(create([]int{1, 2, 4}), create([]int{1, 3, 4}))
	print(result)

	result = mergeTwoLists(create([]int{}), create([]int{0}))
	print(result)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			curr.Next = p2
			p2 = p2.Next
		} else {
			curr.Next = p1
			p1 = p1.Next
		}
		curr = curr.Next
	}

	if p2 != nil {
		curr.Next = p2
	}

	if p1 != nil {
		curr.Next = p1
	}

	return dummy.Next
}

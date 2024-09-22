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
	result := mergeInBetween(create([]int{0,1,2}), 1, 1, create([]int{1000000,1000001,1000002,1000003}))
	print(result)
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	if list1 == nil {
		return nil
	}

	dummy := &ListNode{Next: list1}
	idn := 0
	curr := dummy
	var aNode, bNode *ListNode
	for curr != nil {
		idn++
		curr = curr.Next
		if idn == a {
			aNode = curr
		} else if idn == b + 1 {
			bNode = curr
		}
	}

	aNode.Next = list2
	curr = dummy
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = bNode.Next
	
	return dummy.Next
}

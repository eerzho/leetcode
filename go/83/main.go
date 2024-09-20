package main

import "fmt"

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
	result := deleteDuplicates(create([]int{1, 1, 2}))
	print(result)

	result = deleteDuplicates(create([]int{1, 1, 2, 3, 3}))
	print(result)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	for curr.Next != nil {
		if curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

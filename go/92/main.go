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
	result := reverseBetween(create([]int{1, 2, 3, 4, 5}), 2, 4)
	print(result)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}

	first := dummy
	for i := 0; i < left - 1; i++ {
		first = first.Next
	}

	curr := first.Next
	next := &ListNode{}

	for i := 0; i < right - left; i++ {
		next = curr.Next
		curr.Next = next.Next
		next.Next = first.Next
		first.Next = next
	}

	return dummy.Next
}

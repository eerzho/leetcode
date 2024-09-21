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
	result := removeElements(create([]int{1, 2, 6, 3, 4, 5, 6}), 6)
	print(result)
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy

	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}

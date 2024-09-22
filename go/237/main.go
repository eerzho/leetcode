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
	head := create([]int{4, 5, 1, 9})
	deleteNode(head.Next)
	print(head)
}

func deleteNode(node *ListNode) {
	curr := node 

	var prev *ListNode
	for curr.Next != nil {
		curr.Val = curr.Next.Val
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil
}

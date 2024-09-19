package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func print(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Print("\n")
}

func create(nums []int) *ListNode {
	head := &ListNode{}

	current := head
	for _, num := range nums {
		current.Next = &ListNode{Value: num}
		current = current.Next
	}

	return head.Next
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	list := create(nums)

	print(list)
}

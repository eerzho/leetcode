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
	result := swapPairs(create([]int{1, 2, 3, 4}))
	print(result)
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		first := curr.Next
		second := first.Next

		first.Next = second.Next
		second.Next = first
		curr.Next = second
		curr = first
	}

	return dummy.Next
}

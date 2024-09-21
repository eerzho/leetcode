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
	result := oddEvenList(create([]int{1, 2, 3, 4, 5}))
	print(result)

	// result = oddEvenList(create([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	// print(result)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	odd := fast

	slow := head
	even := slow.Next

	for fast != nil && fast.Next != nil{
		fast.Next = fast.Next.Next
		fast = fast.Next
	}

	print(odd)
	print(even)

	return head
}

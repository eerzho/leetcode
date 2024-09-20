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
	result := removeNthFromEnd(create([]int{1,2,3,4,5}), 2)
	print(result)

	result = removeNthFromEnd(create([]int{1}), 1)
	print(result)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	fast := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := dummy

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
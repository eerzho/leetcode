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
	// case 1
	head := create([]int{1, 3, 4, 7, 1, 2, 6})
	result := deleteMiddle(head)
	print(result)

	head = create([]int{1})
	result = deleteMiddle(head)
	print(result)

	// case 2
	head = create([]int{1, 2, 3, 4})
	result = deleteMiddle(head)
	print(result)
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	fast := head.Next.Next
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

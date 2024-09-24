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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head, Val: -101}
	p1 := dummy
	p2 := head
	prev := dummy

	count := 0
	for p2 != nil {
		if p2.Val == p1.Val {
			p2 = p2.Next
			count++
		} else {
			if count == 0 {
				prev = p1
			} else {
				prev.Next = p2
				count = 0
			}
			p1 = p2
			p2 = p2.Next
		}
	}

	if p1.Next != nil && p1.Val == p1.Next.Val {
		prev.Next = nil
	}

	return dummy.Next
}

func main() {
	result := deleteDuplicates(create([]int{1, 2, 3, 3, 4, 4, 5}))
	print(result)

	result = deleteDuplicates(create([]int{1, 1}))
	print(result)

	result = deleteDuplicates(create([]int{1, 2, 2}))
	print(result)
}

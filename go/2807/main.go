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
	result := insertGreatestCommonDivisors(create([]int{18, 6, 10, 3}))
	print(result)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	for curr.Next != nil {
		a := curr.Val
		b := curr.Next.Val
		for a != b {
			if a > b {
				a = a - b
			} else {
				b = b - a
			}
		}
		tmp := curr.Next
		curr.Next = &ListNode{Next: tmp, Val: a}
		curr = tmp
	}

	return head
}

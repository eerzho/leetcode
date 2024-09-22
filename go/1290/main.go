package main

import (
	"fmt"
	"math"
)

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
	result := getDecimalValue(create([]int{1, 0, 1}))
	fmt.Println(result)
}

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}

	curr := head
	var prev *ListNode

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev

		prev = curr
		curr = tmp
	}

	curr = prev
	var result, raz float64
	for curr != nil {
		result += (math.Pow(2, raz) * float64(curr.Val))
		curr = curr.Next
		raz++
	}

	return int(result)
}

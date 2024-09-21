package main

import "fmt"

type ListNode struct {
	Val int
	Next  *ListNode
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

}

func pairSum(head *ListNode) int {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	curr := slow
	var prev *ListNode
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	
	p1 := prev
	p2 := head
	result := 0
	for p1 != nil {
		sum := p1.Val + p2.Val
		if sum > result {
			result = sum
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	return result
}
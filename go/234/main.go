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
	result := isPalindrome(create([]int{1, 2, 2, 1}))
	fmt.Println(result)

	result = isPalindrome(create([]int{1, 2}))
	fmt.Println(result)
}

func isPalindrome(head *ListNode) bool {
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
	for p1 != nil {
		if p1.Value != p2.Value {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

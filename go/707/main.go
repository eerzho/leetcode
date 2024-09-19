package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}


type MyLinkedList struct {
	Head *Node
	Size int
}

func Constructor() MyLinkedList {
	dummy := Node{}
	return MyLinkedList{
		Head: &dummy,
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}

	current := this.Head
	for i := 0; i < index + 1; i++ {
		current = current.Next
	}

	return current.Value
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}

	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	oldNext := current.Next
	current.Next = &Node{Value: val, Next: oldNext}
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	list := Constructor()
	list.AddAtIndex(0, 1)
	list.AddAtIndex(1, 2)
	list.AddAtIndex(1, 3)
	list.AddAtTail(44)
	list.AddAtHead(4)
	list.DeleteAtIndex(0)

	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))

	current := list.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
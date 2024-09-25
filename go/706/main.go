package main

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type MyHashMap struct {
	cap     int
	buckets []*Node
}

func Constructor() MyHashMap {
	cap := 1433
	buckets := make([]*Node, cap)
	for i := range buckets {
		buckets[i] = &Node{Key: -1}
	}

	return MyHashMap{cap: cap, buckets: buckets}
}

func (this *MyHashMap) Put(key int, value int) {
	hash := this.hash(key)
	head := this.getHead(hash)

	curr := head
	for curr.Next != nil {
		if curr.Key == key {
			curr.Value = value
			return
		}
		curr = curr.Next
	}

	curr.Next = &Node{Key: key, Value: value}
}

func (this *MyHashMap) Get(key int) int {
	hash := this.hash(key)
	head := this.getHead(hash)

	curr := head
	for curr != nil {
		if curr.Key == key {
			return curr.Value
		}
		curr = curr.Next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	hash := this.hash(key)
	head := this.getHead(hash)

	curr := head
	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.cap
}

func (this *MyHashMap) getHead(hash int) *Node {
	return this.buckets[hash]
}

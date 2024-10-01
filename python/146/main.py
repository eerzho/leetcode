from typing import Dict

class Node:
    def __init__(self, key: int, value: int, next: 'Node' = None, prev: 'Node' = None):
        self.key = key
        self.value = value
        self.next = next
        self.prev = prev

class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.data: Dict[int, Node] = {}

        self.head = Node(0, 0, None, None)
        self.tail = Node(0, 0, None, None)

        self.head.next = self.tail
        self.tail.prev = self.head
        
    def remove(self, node: Node) -> None:
        node.prev.next = node.next
        node.next.prev = node.prev

    def addToHead(self, node: Node) -> None:
        node.next = self.head.next
        node.prev = self.head
        self.head.next.prev = node
        self.head.next = node

    def get(self, key: int) -> int:
        if self.data.get(key) == None:
            return -1
        
        node = self.data[key]
        self.remove(node)
        self.addToHead(node)

        return node.value

    def put(self, key: int, value: int) -> None:
        if self.data.get(key) != None:
            node = self.data[key]
            self.remove(node)
            self.addToHead(node)
            node.value = value
            return
        
        if len(self.data) >= self.capacity:
            node = self.tail.prev
            self.remove(node)
            self.data.pop(node.value)            
                
        node = Node(key, value, None, None)
        self.addToHead(node)
        self.data[key] = node


test = str("123")
test.isnumeric


from typing import List, Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if head == None:
            return None
        if k == 0:
            return head

        list_len = 0
        curr = head
        while curr != None:
            curr = curr.next
            list_len += 1

        if list_len == 1:
            return head

        rotate_count = k % list_len
        if rotate_count == 0:
            return head

        curr = head
        count = 0
        while count < list_len - rotate_count - 1:
            curr = curr.next
            count += 1
        if count < list_len - rotate_count - 1 == 0:
            curr = curr.next

        new_head = curr.next
        curr.next = None
        curr = new_head
        while curr.next != None:
            curr = curr.next
        curr.next = head

        return new_head


def create(nums: List[int]) -> ListNode:
    head = ListNode()

    curr = head
    for num in nums:
        curr.next = ListNode(val=num)
        curr = curr.next

    return head.next

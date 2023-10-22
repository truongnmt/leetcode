# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        less_part = ListNode()
        greater_part = ListNode()

        less_iter = less_part
        greater_iter = greater_part

        while head:
            if head.val < x:
                less_iter.next = head
                less_iter = head
            else:
                greater_iter.next = head
                greater_iter = head
            head = head.next
            
        greater_iter.next = None
        less_iter.next = greater_part.next

        return less_part.next

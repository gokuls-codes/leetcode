# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def splitListToParts(self, head: Optional[ListNode], k: int) -> List[Optional[ListNode]]:
        temp = head
        len = 0
        while temp != None:
            temp = temp.next
            len += 1

        count = len // k
        odd = len % k

        ans = []
        temp = head

        for i in range(k):
            t2 = temp
            req = count
            if i < odd:
                req += 1
            for j in range(req-1):
                t2 = t2.next
            t3 = None
            if t2:
                t3 = t2.next
                t2.next = None
            ans.append(temp)
            temp = t3

        return ans
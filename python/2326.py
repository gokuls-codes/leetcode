# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def spiralMatrix(self, m: int, n: int, head: Optional[ListNode]) -> List[List[int]]:
        ans = [[-1 for _ in range(n)] for _ in range(m)]

        t, b = 0, m - 1
        l, r = 0, n - 1

        while head:
            for j in range(l, r + 1):
                if not head: break
                ans[t][j] = head.val
                head = head.next
            
            for i in range(t + 1, b + 1):
                if not head: break
                ans[i][r] = head.val
                head = head.next
            
            for j in range(r - 1, l - 1, -1):
                if not head: break
                ans[b][j] = head.val
                head = head.next

            for i in range(b - 1, t, -1):
                if not head: break
                ans[i][l] = head.val
                head = head.next

            t += 1
            b -= 1
            l += 1
            r -= 1

        return ans
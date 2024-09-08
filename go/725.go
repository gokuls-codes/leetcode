/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func splitListToParts(head *ListNode, k int) []*ListNode {
    temp := head
    len := 0
    for temp != nil {
        temp = temp.Next
        len++
    }
    count := len / k
    odd := len % k
    
    ans := []*ListNode{}
    temp = head

    for i := range(k) {
        t2 := temp
        req := count
        if i < odd {
            req++
        }
        for _ = range(req - 1) {
            t2 = t2.Next
        }
        var t3 *ListNode = nil
        if t2 != nil {
            t3 = t2.Next
            t2.Next = nil
        }
        ans = append(ans, temp)
        temp = t3
    }

    return ans
}
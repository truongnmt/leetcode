/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    tmp := head
    newPrev := tmp.Next
    oldNext := newPrev.Next
    newPrev.Next = tmp
    tmp.Next = nil
    
    for true {
        tmp = newPrev
        newPrev = oldNext
        if newPrev == nil {
            return tmp
        }
        oldNext = newPrev.Next
        newPrev.Next = tmp
    }
    return nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    
    m := make(map[*ListNode]int)
    for head != nil {
        if m[head] == 1 {
            return head
        }
        m[head] = 1
        head = head.Next
    }
    return nil
}

func detectCycle2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    
    slow := head
    fast := head
    var meetingPoint *ListNode
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            meetingPoint = slow
            break
        }
    }
    
    if meetingPoint == nil {
        return nil
    }
    for head != meetingPoint {
        head = head.Next
        meetingPoint = meetingPoint.Next
    }
    return head
}

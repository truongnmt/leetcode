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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    for tmp:=head; tmp != nil; tmp=tmp.Next {
        for tmp.Next != nil && tmp.Val == tmp.Next.Val {
            tmp.Next = tmp.Next.Next
        }
    }
    return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 || (lists[0] == nil && len(lists) == 1) {
        return nil
    }
    
    result := lists[0]
    for _, list := range lists[1:len(lists)] {
        result = mergeTwoLists(result, list)
    }
    
    return result
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    
    head := l1
    if l2.Val < l1.Val {
        head = l2
        l2 = l2.Next
    } else {
        l1 = l1.Next
    }
    
    tmp := head
    for l1 != nil || l2 != nil {
        if l1 == nil {
            tmp.Next = l2
            break;
        }
        if l2 == nil {
            tmp.Next = l1
            break;
        }
        
        if l1.Val < l2.Val {
            tmp.Next = l1
            tmp = l1
            l1 = l1.Next
        } else {
            tmp.Next = l2
            tmp = l2
            l2 = l2.Next
        }
    }
    
    return head
}

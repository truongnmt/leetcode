/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    var arrNode []*ListNode
    dummy := head
    
    for dummy != nil {
        arrNode = append(arrNode, dummy)
        dummy = dummy.Next
    }
    
    // reverse k element at a time
    for i:=0;i<len(arrNode)-k+1;i+=k {
        for j,k:=i,i+k-1; j<k; j,k=j+1,k-1 {
            arrNode[j], arrNode[k] = arrNode[k], arrNode[j]
        }
    }
    
    for i, node := range arrNode {
        if i == len(arrNode)-1 {
            node.Next = nil
        } else {
            node.Next = arrNode[i+1]
        }
    }
    
    return arrNode[0]
}

// ==============================

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{0, head}
    groupPrev := dummy
    
    for true {
        kth := getKthNode(groupPrev, k)
        if kth == nil {
            break
        }
        
        groupNext := kth.Next
        
        // reverse node in the group
        prev, cur := groupNext, groupPrev.Next
        
        for cur != groupNext {
            tmp := cur.Next
            cur.Next = prev
            prev = cur
            cur = tmp
        }
        
        tmp := groupPrev.Next
        groupPrev.Next = kth
        groupPrev = tmp
    }
    
    return dummy.Next
}

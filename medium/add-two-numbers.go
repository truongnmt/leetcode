/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// extra = 0
// while l1 || l2 != null
// if l1 == null || l2.Val == null, l1.Val = 0
// sum = l1+l2+extra
// tmp = new node
// dummy = tmp
// dummy.Val = sum%10
// if sum > 9, extra = l1+l2-9

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	extra := 0
	result := &ListNode{}
	dummy := result

	for l1 != nil || l2 != nil {
		p, q := 0, 0
		if l1 == nil {
			p = 0
		} else {
			p = l1.Val
		}
		if l2 == nil {
			q = 0
		} else {
			q = l2.Val
		}

		sum := p + q + extra
		extra = sum / 10
		// fmt.Printf("sum=%v p=%v q=%v extra=%v\n", sum, p, q, extra)

		dummy.Val = sum % 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			if extra == 1 {
				dummy.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
			} else {
				dummy.Next = nil
			}
		} else {
			dummy.Next = &ListNode{}
		}
		dummy = dummy.Next
	}

	return result
}

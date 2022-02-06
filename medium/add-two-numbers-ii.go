/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stackL1 := []int{}
	for l1 != nil {
		stackL1 = append(stackL1, l1.Val)
		l1 = l1.Next
	}

	stackL2 := []int{}
	for l2 != nil {
		stackL2 = append(stackL2, l2.Val)
		l2 = l2.Next
	}

	var result *ListNode
	sum := 0

	for len(stackL1) > 0 || len(stackL2) > 0 || sum > 0 {
		left, right := 0, 0
		if len(stackL1) > 0 {
			left = stackL1[len(stackL1)-1]
			stackL1 = stackL1[:len(stackL1)-1]
		}
		if len(stackL2) > 0 {
			right = stackL2[len(stackL2)-1]
			stackL2 = stackL2[:len(stackL2)-1]
		}
		sum += left + right

		dummy := &ListNode{
			Val:  sum % 10,
			Next: result,
		}
		result = dummy

		sum /= 10
	}

	return result
}

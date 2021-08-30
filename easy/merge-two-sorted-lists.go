package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{}}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{}}}}

	result := mergeTwoLists(l1, l2)

	for i := result; i.Next != nil; i = i.Next {
		fmt.Print(i.Val)
	}

	// fmt.Println(result)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	myMap := make(map[int]int)
	keys := make([]int, 0)

	for listTemp := l1; listTemp != nil; listTemp = listTemp.Next {
		myMap[listTemp.Val] += 1
	}
	for listTemp := l2; listTemp != nil; listTemp = listTemp.Next {
		myMap[listTemp.Val] += 1
	}
	for k, _ := range myMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	listTemp := &ListNode{}
	listPrev := &listTemp
	*listPrev = nil
	for i := len(keys) - 1; i >= 0; i-- {
		for j := 0; j < myMap[keys[i]]; j++ {
			if listPrev == nil {
				listTemp = &ListNode{Val: keys[i]}
			} else {
				listTemp = &ListNode{Val: keys[i], Next: *listPrev}
			}

			listPrev = &listTemp
		}
	}

	return listTemp
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
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

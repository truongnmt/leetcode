// Given a sorted array A[] with possibly duplicate elements, write a program to find indexes of first and last occurrences of an element k in the given array.

// Problem Note

// The algorithm’s runtime complexity must be in the order of O(log n).
// If the target is not found in the array, return [-1, -1].
// Example 1

// Input: A[] = [1, 3, 5, 5, 5, 5 ,28, 37, 42], target = 5
// Output: [2, 5]
// Explanation: First Occurrence = 2 and Last Occurrence = 5
// Example 2

// Input: A[] = [1, 3, 5, 5, 5, 5 ,7, 28, 37], target = 7
// Output: [6, 6]
// Example 3

// Input: A[] = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]

// ---

// Linear Search: O(N)
// Binary Search: do it twice with return left and return right O(logN)

// return l: l, r = 0, n
// return r: l, r = -1, n-1

package main

import (
	"fmt"
)

func main() {
	A1 := []int{1, 3, 5, 5, 5, 5 ,28, 37, 42}
	A2 := []int{1, 3, 5, 5, 5, 5 ,7, 28, 37}
	A3 := []int{5, 7, 7, 8, 8, 10 }

	fmt.Println(searchRange(A1, 5))
	fmt.Println(searchRange(A2, 7))
	fmt.Println(searchRange(A3, 6))
}

func searchRange(A []int, k int) []int {
	ans := []int{-1,-1}

	// get ans[0]
	left, right := 0, len(A) // 0 9
	for left+1 < right {
		// fmt.Printf("left=%v right=%v\n", left, right)
		middle := (left+right)/2 // 4 2 1
		if A[middle] <= k {
			left = middle // right = 4 2
		} else {
			right = middle // left = 1
		}
	}
	// fmt.Println(left)
	if A[left] == k { ans[1] = left }
	
	left, right = -1, len(A)-1
	for left+1 < right {
		// fmt.Printf("left=%v right=%v\n", left, right)
		middle := (left+right)/2
		if A[middle] < k {
			left = middle
		} else {
			right = middle
		}
	}
	// fmt.Println(right)
	if A[right] == k { ans[0] = right }

	return ans
}


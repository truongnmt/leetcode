package main

import (
	"fmt"
)

func main() {
	// arr := []int{1, 2, 4, 5, 6, 6, 8, 9}
	// arr2 := []int{2, 5, 6, 7, 8, 8, 9}
	arr3 := []int{-12, -5, 6, 7, 8, 8, 9}
	

	// fmt.Println(findClosestNumberArray(arr, 11))
	// fmt.Println(findClosestNumberArray(arr2, 4))
	fmt.Println(findClosestNumberArray(arr3, 0))
}

func findClosestNumberArray(arr []int, target int) int {

	if target > arr[len(arr)-1] {
		return arr[len(arr)-1]
	} else if target < arr[0] {
		return arr[0]
	}

	l, r := 0, len(arr)
	for l+1<r {
		m := (l+r)/2
		if arr[m] <= target {
			l = m
		} else {
			r = m
		}
	}

	fmt.Printf("l=%v r=%v\n", l, r)

	if abs(target-arr[l]) < abs(target-arr[r]) {
		return arr[l]
	} else {
		return arr[r]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
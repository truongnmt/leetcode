package main

import (
	"fmt"
)

func main() {
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}

	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}

func partition(arr []int, low, high int) int {
	p := arr[high]
	for j := low; j < high; j++ {
		if arr[j] < p {
			arr[low], arr[j] = arr[j], arr[low]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	return low
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

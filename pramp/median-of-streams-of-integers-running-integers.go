// 1 2 3 4 | 5 6 7 8
// max heap: 4 3 2 1
// min heap: 5 6 7 8
// max heap left, min heap right

// max heap | min heap  
// 1 2 3  4 | 5 6 7 8
// addNum: O<logN>
// if num is larger than min heap, put value in min heap  O(logN)
// else put in max heap
// if len of maxH > min heap+1
//   pop maxHeap push minH                          O<logN> + O<logN>
// else if len of minHeap > maxHeap + 1
//   pop minHeap push maxHeap   
// findMedian
// if len maxH > len minH return get MaxH
// else len maxH < len minH return get MinH
// else return (getMax + getMin) /2 


package main

import (
	"fmt"
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]

	return x
}

func main() {
	A := []int{5, 15, 1, 3, 2, 8, 7, 9, 10, 6, 11, 4}
	// A2 := []int{5, 10, 15}

	maxHeap := &IntHeap{}
	minHeap := &IntHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)

	// 9
	// max heap | min Heap
	// 1 2 3 5 | 6 7 8 9
	for _, num := range A {
		if minHeap.Len() == 0 || num > (*minHeap)[0] {
			heap.Push(minHeap, num)
		} else {
			heap.Push(maxHeap, -num)
		}

		if maxHeap.Len() > minHeap.Len() + 1 {
			heap.Push(minHeap, -heap.Pop(maxHeap).(int))
		} else if minHeap.Len() > maxHeap.Len() + 1 {
			heap.Push(maxHeap, -heap.Pop(minHeap).(int))
		}

		if maxHeap.Len() < minHeap.Len() {
			fmt.Println(float64((*minHeap)[0]))
		} else if maxHeap.Len() > minHeap.Len() {
			fmt.Println(float64(-(*maxHeap)[0]))
		} else {
			fmt.Println(float64((-(*maxHeap)[0] + (*minHeap)[0]))/2)
		}		
	}
}

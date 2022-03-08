// max heap | min heap  
// 1 2 3  4 | 5 6 7 8

// addNum: O<logN>
// if num is larger than min heap, put value in min heap  O(logN)
// else put in max heap
// if len of maxH > min heap+1
//   pop maxHeap push minH                          O<logN> + O<logN>
// else maxH < minHeap + 1
//   pop minHeap push maxHeap                       O<logN> + O<logN>

// findMedian
// if len maxH > len minH return get MaxH
// else len maxH < len minH return get MinH
// else return (getMax + getMin) /2

// import "container/heap"

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


type MedianFinder struct {
    maxHeap, minHeap *IntHeap
}


func Constructor() MedianFinder {
    maxHeap := &IntHeap{}
    minHeap := &IntHeap{}
    heap.Init(maxHeap)
    heap.Init(minHeap)
    return MedianFinder{maxHeap, minHeap}
}

// max heap | min heap  
// 1 2 3  4 | 5 6 7 8
// addNum: O<logN>
// if num is larger than min heap, put value in min heap  O(logN)
// else put in max heap
// if len of maxH > min heap+1
//   pop maxHeap push minH                          O<logN> + O<logN>
// else if len of minHeap > maxHeap + 1
//   pop minHeap push maxHeap                       O<logN> + O<logN>
func (this *MedianFinder) AddNum(num int)  {
    if this.minHeap.Len() == 0 || num > (*this.minHeap)[0] {
        heap.Push(this.minHeap, num)
    } else {
        heap.Push(this.maxHeap, -num)
    }
    
    if this.maxHeap.Len() > this.minHeap.Len() + 1 {
        heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
    } else if this.minHeap.Len() > this.maxHeap.Len() + 1 {
        heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
    }
}

// findMedian
// if len maxH > len minH return get MaxH
// else len maxH < len minH return get MinH
// else return (getMax + getMin) /2
func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() > this.minHeap.Len() {
        return float64(-(*this.maxHeap)[0])
    } else if this.maxHeap.Len() < this.minHeap.Len() {
        return float64((*this.minHeap)[0])
    }
    
    return float64(-(*this.maxHeap)[0] + (*this.minHeap)[0])/2
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
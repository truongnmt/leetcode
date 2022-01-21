import (
    "container/heap"
)

func distance(point []int) int {
    return point[0]*point[0]+point[1]*point[1]
}

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return distance(pq[i]) > distance(pq[j])
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    point := old[n-1]
    *pq = old[0:n-1]
    return point
}

func kClosest(points [][]int, k int) [][]int {
    //if len(heap) < k
    //  heappush
    // else if distance(point) < distance(heap[0]) 
    //  heap pop then push 
    
    h := &PriorityQueue{}
    heap.Init(h)
    
    for _, point := range points {
        if h.Len() < k {
            heap.Push(h, point)
        } else {
            if distance(point) < distance((*h)[0]) {
                heap.Pop(h)
                heap.Push(h, point)
            }
        }
    }
    

    var arr [][]int
    for _, point := range *h {
        arr = append(arr, point)
    }
    return arr
}

type MyCircularQueue struct {
    size int
    q []int
}


func Constructor(k int) MyCircularQueue {
    var cq = new(MyCircularQueue)
    cq.size = k
    cq.q = []int{}
    return *cq
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if len(this.q) >= this.size { return false }
    
    this.q = append(this.q, value) // [1 2 3] -> [2 3 4]
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if len(this.q) == 0 { return false }
    
    this.q = this.q[1:] // [2 3]
    return true
}


func (this *MyCircularQueue) Front() int {
    if len(this.q) == 0 { return -1 }
    
    return this.q[0]
}


func (this *MyCircularQueue) Rear() int {
    if len(this.q) == 0 { return -1 }
    
    return this.q[len(this.q)-1] // 3 -> 4
}


func (this *MyCircularQueue) IsEmpty() bool {
    if len(this.q) == 0 { return true }
    
    return false
}


func (this *MyCircularQueue) IsFull() bool {
    if len(this.q) == this.size { return true }
    
    return false
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
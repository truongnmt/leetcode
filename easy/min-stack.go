type MinStack struct {
    stack []int
    minStack []int
}


// Init with 2 stack
func Constructor() MinStack {
    var minStack = new(MinStack)
    minStack.stack = []int{}
    minStack.minStack = []int{}
    return *minStack
}

// Push to stack
// Push to minStack, if new element smaller than last element in minStack then push
func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    
    if len(this.minStack) == 0 {
        this.minStack = append(this.minStack, val)
    } else if this.minStack[len(this.minStack)-1] >= val {
        this.minStack = append(this.minStack, val)
    }
}

// Pop the last element in stack
// Pop in minStack, if last element in stack equal to last element in minStack then pop in minStack as well
func (this *MinStack) Pop()  {
    nStack:=len(this.stack)-1
    nMinStack:=len(this.minStack)-1
    if this.stack[nStack] == this.minStack[nMinStack] {
        this.minStack = this.minStack[:(len(this.minStack)-1)]
    }
    
    this.stack = this.stack[:(len(this.stack)-1)]
}

// Return last element in stack
func (this *MinStack) Top() int {
    nStack:=len(this.stack)-1
    return this.stack[nStack]
}

// Return last element in minStack
func (this *MinStack) GetMin() int {
    return this.minStack[(len(this.minStack)-1)]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

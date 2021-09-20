var kV = make(map[int]int)
var kP = make(map[int]*LRUCache)

type ListNode struct {
    Key int
    Next *ListNode
    Prev *ListNode
}

type LRUCache struct {
    kV map[int]int
    kP map[int]*ListNode
    top *ListNode
    bot *ListNode
    capacity int
}


func Constructor(capacity int) LRUCache {
    var lruC = new(LRUCache)
    lruC.capacity = capacity
    lruC.kV = make(map[int]int)
    lruC.kP = make(map[int]*ListNode)
    return *lruC
}

func printList(this *LRUCache) {
    node := this.top
    for node != nil {
        fmt.Printf("%v -> ", node.Key)
        node = node.Next
    }
    fmt.Println()
    fmt.Println(this.kV)
    fmt.Println()
}

func (this *LRUCache) Get(key int) int {
    // fmt.Printf("Get key: %v\n", key)
    if node, ok := this.kP[key]; ok {
        if node != this.top && node != this.bot && node.Prev != nil && node.Next != nil {
            node.Prev.Next = node.Next
            node.Next.Prev = node.Prev
            
            node.Prev = this.bot
            this.bot.Next = node
            this.bot = node
            this.bot.Next = nil
        } else if this.capacity == 1 {
            return this.kV[key]
        } else if node == this.top && node.Next != nil {
            this.top = this.top.Next
            this.top.Prev = nil
            
            this.bot.Next = node
            node.Prev = this.bot
            node.Next = nil
            this.bot = node
        }
        // fmt.Printf("Return: %v\n",this.kV[key])
        // printList(this)
        return this.kV[key]
    }
        
    // fmt.Printf("Return: -1\n")
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    // fmt.Printf("Put k-v: %v-%v\n", key, value)
    if _, ok := this.kV[key]; ok {
        this.kV[key] = value
        _ = this.Get(key)
    } else {
        this.kV[key] = value
        node := new(ListNode)
        node.Key = key
        if this.top == nil {
            this.top = node
            this.bot = node
        } else {
            node.Prev = this.bot
            this.bot.Next = node
            this.bot = node
        }
        this.kP[key] = node

        // fmt.Printf("len(this.kV): %v, this.capacity: %v\n", len(this.kV), this.capacity)
        if len(this.kV) > this.capacity {
            delete(this.kV, this.top.Key)
            delete(this.kP, this.top.Key)
            this.top = this.top.Next
        }
        // fmt.Println(this.kV)
    }
    // printList(this)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

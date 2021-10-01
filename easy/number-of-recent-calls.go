type RecentCounter struct {
    st []int
}


func Constructor() RecentCounter {
    var recentCounter = new(RecentCounter)
    recentCounter.st = []int{}
    return *recentCounter
}


func (this *RecentCounter) Ping(t int) int {
    for _, ping := range(this.st) {
        if ping < (t-3000) {
            this.st = this.st[1:]
        }
    }
    this.st = append(this.st, t)
    return len(this.st)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
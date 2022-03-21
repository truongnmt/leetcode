func jump(nums []int) int {
    count := 0
    l, r := 0, 0
    
    for r < len(nums)-1 {
        farthest := 0
        for i:=l;i<=r;i++ {
            if i+nums[i]>farthest {
                farthest = i+nums[i]
            }
        }
        l, r = r+1, farthest
        count += 1
    }
    
    return count
}

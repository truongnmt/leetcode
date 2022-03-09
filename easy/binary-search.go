// -1,0,3,5,9,12
// len = 6

func search(nums []int, target int) int {
    l, r := 0, len(nums) // 4 5
    
    for l+1<r {
        m := (l+r)/2     // 3 4 5
        if nums[m] <= target {
            l = m        
        } else {
            r = m
        }
    }
    
    if nums[l] != target {
        return -1
    }
    
    return l
}

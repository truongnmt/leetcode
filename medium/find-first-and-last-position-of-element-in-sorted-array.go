// nums=5,7,7,8,8,10    target=8
//      0 1 2 3 4 5  6
//          l r
// output=3,4

func searchRange(nums []int, target int) []int {
    ans := []int{-1,-1}
    if len(nums) == 0 {
        return ans
    }
    
    // binary search return left to get last position
    // TC O(log n)
    l, r := 0, len(nums)
    for l+1<r {
        m := (l+r)/2
        if nums[m] <= target {
            l = m
        } else {
            r = m
        }
    }
    if nums[l] == target {
        ans[1] = l
    }    
    
    // binary search return right to get first position
    // TC O(log n)
    l, r = -1, len(nums)-1
    for l+1 < r {
        m := (l+r)/2
        if nums[m] < target {
            l = m
        } else {
            r = m
        }
    }
    if nums[r] == target {
        ans[0] = r
    }
    
    return ans
}

func longestConsecutive(nums []int) int {
    if len(nums) < 1 {
        return 0
    }
    
    mapNum := make(map[int]bool, len(nums))
    for _, num := range nums {
        mapNum[num] = true
    }
    
    maxCount := 0
    
    for num := range mapNum {
        if !mapNum[num-1] {
            curCount := 0
            for mapNum[num+curCount] {
                curCount += 1   
            }
            if curCount > maxCount { maxCount = curCount }
        }
    }
    
    return maxCount
}

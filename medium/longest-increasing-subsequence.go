func lengthOfLIS(nums []int) int {
    n := len(nums)
    d := make([]int, n)
    max := 1
    
    for i:=1;i<n;i++ {
        for j:=0;j<i;j++ {
            if d[j]==0 {
                d[j]=1
            }
            if nums[j] < nums[i] {
                if d[j]+1 > d[i] {
                    d[i] = d[j]+1
                }
            }
        }
        if d[i] > max {
            max = d[i]
        }
    }
    
    return max
}

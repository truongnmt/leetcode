//https://play.golang.org/p/CVNce1o_6A-

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return nil
    }
    
    sort.Ints(nums)
    hashAns := make(map[[3]int]bool)
    for i, num := range nums {
        j, k := i+1, len(nums)-1
        for j<k {
            // fmt.Printf("i=%v j=%v k=%v\n", num, nums[j], nums[k])
            sum := num+nums[j]+nums[k]
            if sum > 0 {
                k -= 1
            } else if sum == 0 {
                // ans = append(hashAns, []int{num, nums[j], nums[k]})
                hashAns[[3]int{num, nums[j], nums[k]}] = true
                j += 1
            } else if sum < 0 {
                j += 1
            }
        }
    }
    
    ans := make([][]int, 0)
    for v, _ := range hashAns {
        ans = append(ans, []int{v[0], v[1], v[2]})
    }

    return ans
}

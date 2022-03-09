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

func threeSum(nums []int) [][]int {
    
    //Input: nums = [-1,0,1,2,-1,-4]
    //Output: [[-1,-1,2],[-1,0,1]]
    // contain no duplicate triplets
    
//     sort nums to -4 -1 -1 0 1 2
//                      i    l r  
//     loop num in nums
//     left, right = num+1, len(num)-1
//     while 0<left<right<len(num)
//       if num+left+right==0 append to answer
//       else if num+left+right < 0 left++
//       else if num+left+right > 0 right--

//     return answer
    
    if len(nums) < 2 {
        return [][]int{}
    }
    
    sort.Ints(nums)
    fmt.Println(nums)
    ans := make([][]int, 0)
    hashAns := make(map[[3]int]bool)
    
    for i:=0;i<len(nums)-2;i++ {
        left, right := i+1, len(nums)-1
        // fmt.Printf("i=%v left=%v right=%v\n", i, left, right)
        if i>1 && nums[i]==nums[i-1] {
            continue
        }
        for left<right {
            sum := nums[i]+nums[left]+nums[right] 
            if sum==0 {
                hashAns[[3]int{nums[i], nums[left], nums[right]}] = true
                // m = append(m, []int{nums[i], nums[left], nums[right]})
                left++
                right--
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }
    
    for arr, _ := range hashAns {
        ans = append(ans, []int{arr[0], arr[1], arr[2]})
    }
    
    return ans
}
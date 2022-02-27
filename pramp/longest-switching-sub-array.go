// An array is called "switching" if the odd and even elements are equal.
// Example:
// [2,4,2,4] is a switching array because the members in even positions (indexes 0 and 2) and odd positions (indexes 1 and 3) are equal.

// If A = [3,7,3,7, 2,1,2], the switching sub-arrays are:
// == > [3,7,3,7] and [2,1,2]
// Therefore, the longest switching sub-array is [3,7,3,7] with length = 4.

// As another example if A = [1,5,6,0,1,0], the the only switching sub-array is [0,1,0].
// Another example: A= [7,-5,-5,-5,7,-1,7], the switching sub-arrays are [7,-1,7] and [-5,-5,-5].

// Question:
// Write a function that receives an array and find its longest switching sub-array.

// [3,7,3,7, 2,1,2]
//  l     r 
//  0 1 2 3  4 5 6

// left = 0
// loop right=2;right<len(nums);right++
// valid = compare right == right-2 
// if valid 
//   if not switchingSubArrayDetected
//     left = right - 2
//   max = max(max, right - left)
//   switchingSubArrayDetected = true
// else 
//   switchingSubArrayDetected = false


func main() {   //0 1 2 3 4 5
    //            l   r
    // nums := []int{1,5,6,0,1,0} // [3,5] 3
    // nums := []int{7,-5,-5,-5,7,-1,7} // [1,3] 3
    nums := []int{3,7,3,7,2,1,2} // [0, 2] 4
    fmt.Println(longestSwitchingSubArray(nums))
}

func longestSwitchingSubArray(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    
    left := 0
    max := 0    
    switchingSubArrayDetected := false
    for right:=2;right<len(nums);right++ {
        valid := false
        if nums[right] == nums[right-2] { valid = true }
                
        if valid {
            if !switchingSubArrayDetected { left = right - 2 }
            if (right-left+1) > max { 
                max = right-left+1 
                fmt.Printf("left=%v right=%v\n", left, right)
            }
        } else {
            switchingSubArrayDetected = false
        }
    }
    
    
    return max
}





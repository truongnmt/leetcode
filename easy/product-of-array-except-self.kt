// [1,2,3,4]
//  i
//  
// [1, 1, 2,6]: prefix
// [24 12 4 1]: postfix
// [24 12 8 6]: result

class Solution {
    fun productExceptSelf(nums: IntArray): IntArray {
        var result = IntArray(nums.size) {1}
        var lastProduct = 1

        nums.forEachIndexed { i, num ->
            if (i>0) result[i] = lastProduct * nums[i-1]
            lastProduct = result[i]
        }
        
        lastProduct = 1
        for (i in (nums.size-1) downTo 0) {
            if (i<nums.size-1) lastProduct = lastProduct * nums[i+1]
            result[i] *= lastProduct
        }

        return result


        // var prefix = IntArray(nums.size) {nums[0]}

        // nums.forEachIndexed { i, num -> 
        //     if (i>0) prefix[i] = prefix[i-1]*num
        // }

        // var postfix = IntArray(nums.size) {nums[nums.size-1]}

        // for (i in (nums.size-1) downTo 0) {
        //     if (i<(nums.size-1)) postfix[i] = postfix[i+1]*nums[i]
        // }

        // var result = IntArray(nums.size) {1}

        // result.forEachIndexed { i, num -> 
        //     if (i==0) result[0] = postfix[1]
        //     else if (i==nums.size-1) result[i] = prefix[i-1]
        //     else result[i] = prefix[i-1]*postfix[i+1]
        // }

        // return result
    }
}
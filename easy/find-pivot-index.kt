//        i
// [1,7,3,6,5,6]

// rightSum=28,27,20,17,11
// leftSum=0,1,8,11
// for i in 0..(nums.size)
// rightSum=-nums[i]
// leftSum=+nums[i-1]
// if leftSum == rightSum ? return i


class Solution {
    fun pivotIndex(nums: IntArray): Int {
        var rightSum = 0
        var leftSum = 0

        for (num in nums) rightSum += num

        for (i in 0 until nums.size) {
            rightSum -= nums[i]
            if (i>0) leftSum += nums[i-1]
            if (leftSum == rightSum) return i
        }
        
        return -1
    }
}
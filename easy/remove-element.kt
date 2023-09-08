// Intuition
//
// define an index
// loop through nums
// nums[i]!=val ? nums[index] = nums[i]; index++
// return index

class Solution {
    fun removeElement(nums: IntArray, `val`: Int): Int {
        // var j = nums.size-1

        // for (i in nums.indices) {
        //     while (nums[j]==`val` && j>0) {
        //         j--
        //     }
            
        //     if (nums[i]==`val` && i<j) {
        //         nums[i] = nums[j]
        //         j--
        //     }
        // }
        //return j+1

        // var index = 0
        // for (i in nums.indices) {
        //     if (nums[i]!=`val`) {
        //         nums[index] = nums[i]
        //         index++
        //     }
        // }
        // return index

        var counter = 0
        nums.forEach { if (it!=`val`) nums[counter++] = it}
        return counter
    }
}
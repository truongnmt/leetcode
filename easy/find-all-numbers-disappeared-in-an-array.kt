// Intution
// 
// since there will be a duplicate number, if we use the number as index, it will point to the same place
// we can create theee side effect by mark number as negative in the index
// filter the list for positive number

import kotlin.math.abs

class Solution {
    fun findDisappearedNumbers(nums: IntArray): List<Int> {
        val result = mutableListOf<Int>()
        
        // nums.map {
        //     val absIdx = abs(it)-1
        //     if (nums[absIdx]>0) nums[absIdx] = -nums[absIdx]
        // }
        // this map method use more memory

        for (i in nums) {
            val absIdx = abs(i)-1
            if (nums[absIdx]>0) nums[absIdx] = -nums[absIdx]
        }

        nums.forEachIndexed { i, num ->
            if (num>0) {
                result.add(i+1)
            }
        }

        return result
    }
}


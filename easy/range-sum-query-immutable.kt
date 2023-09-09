// NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
// [-2, 0, 3, -5, 2, -1]
//  0  -2  1  -4 -2  -3
// [0, 2]: -2+0+3 = 1 = 1 - 0
// [2, 5]: 3-5+2-1=-1 = -3 -  -2 = -1  
//
// numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
// numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
// numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

class NumArray(nums: IntArray) {
    private val prefixSum = IntArray(nums.size)
    private val numsArr = nums

    init {
        prefixSum[0] = nums[0]
        for (i in 1 until nums.size) {
            prefixSum[i] = prefixSum[i-1] + nums[i]
        }
    }

    fun sumRange(left: Int, right: Int): Int {
        if (left==right) return numsArr[left]
        if (left==0 && right!=0) return prefixSum[right]

        return prefixSum[right] - prefixSum[left-1]
    }
}

// class NumArray(val nums: IntArray) {
//     private val prefixSum = IntArray(nums.size+1) {0}
//     init{
//         for (i in 1..nums.size) {
//             prefixSum[i] = prefixSum[i-1] + nums[i-1]
//         }
//     }
//     fun sumRange(left: Int, right: Int): Int {
//         return prefixSum[right+1] - prefixSum[left]
//     }

// }

/**
 * Your NumArray object will be instantiated and called as such:
 * var obj = NumArray(nums)
 * var param_1 = obj.sumRange(left,right)
 */
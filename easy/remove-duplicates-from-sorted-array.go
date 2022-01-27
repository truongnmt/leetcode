func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	uniqueIndex := 1
	for i, _ := range nums {
		if i+1 < len(nums) && nums[i] != nums[i+1] {
			nums[uniqueIndex] = nums[i+1]
			uniqueIndex++
		}

	}

	return uniqueIndex
}

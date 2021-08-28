func containsDuplicate(nums []int) bool {
	quickSort(nums, 0, len(nums)-1)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			return true
		}
	}
	return false
}

func partition(arr []int, low, high int) int {
	p := arr[high]
	for j := low; j < high; j++ {
		if arr[j] < p {
			arr[j], arr[low] = arr[low], arr[j]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	return low
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

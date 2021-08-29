import "fmt"

func topKFrequent(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	m := make(map[int]int)
	invM := make(map[int][]int)
	var freq []int
	i := 0

	for j, num := range nums {
		i++
		if j == len(nums)-1 || (j < len(nums)-1 && num != nums[j+1]) {
			m[num] = i
			invM[i] = append(invM[i], num)
			i = 0
		}
	}
	for i := range invM {
		freq = append(freq, i)
	}

	quickSort(freq, 0, len(freq)-1)
	var result []int
	for i = 0; i < len(freq); i++ {
		fmt.Println(invM[freq[len(freq)-i-1]])
		if len(result) >= k {
			break
		}
		result = append(result, invM[freq[len(freq)-i-1]]...)
	}

	return result
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

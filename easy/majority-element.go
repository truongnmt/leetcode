func majorityElement(nums []int) int {
    quickSort(nums, 0, len(nums)-1)
    return nums[(len(nums))/2]
}

func partition(arr []int, low, high int) int {
    p := arr[high]
    for j:=low;j<high;j++ {
        if arr[j] < p {
            arr[low], arr[j] = arr[j], arr[low]
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

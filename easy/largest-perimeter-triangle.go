func largestPerimeter(nums []int) int {
    quickSort(nums, 0, len(nums)-1)
    for i:=len(nums)-1;i>1;i-- {
        if nums[i] < nums[i-1]+nums[i-2] {
            return nums[i]+nums[i-1]+nums[i-2]
        }
    }
    return 0
}

func partition(arr []int, low, high int) int {
    p := arr[high]
    for j:=low;j<high;j++ {
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

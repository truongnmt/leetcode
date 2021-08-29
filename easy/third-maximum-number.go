func thirdMax(nums []int) int {
    quickSort(nums, 0, len(nums)-1)
    
    if len(nums)<3 {
        return nums[len(nums)-1]
    }
    
    i:=1
    for j:=len(nums)-1;j>0;j-- {
        if i<3 && nums[j] != nums[j-1] {
            i++
        }
        if i==3 {
            return nums[j-1]
        }
        
    }
    return nums[len(nums)-1]
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

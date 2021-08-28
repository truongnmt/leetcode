// Runtime: 88 ms, faster than 51.92% of Go online submissions for Rank Transform of an Array.
// Memory Usage: 12.2 MB, less than 73.08% of Go online submissions for Rank Transform of an Array.


func arrayRankTransform(arr []int) []int {
    cloneArr := make([]int, len(arr))
    copy(cloneArr, arr)
    quickSort(cloneArr, 0, len(cloneArr)-1)
    
    m := make(map[int]int)
    i := 1
    for j, num := range(cloneArr) {
        m[num] = i
        if j<len(cloneArr)-1 && num != cloneArr[j+1] {
            i++
        }
    }
    
    for i=0; i< len(arr);i++ {
        arr[i] = m[arr[i]]
    }
    return arr
}

func partition(arr []int, low, high int) int {
    p := arr[high]
    for j:=low; j<high; j++ {
        if arr[j] < p {
            arr[low], arr[j] = arr[j], arr[low]
            low++
        }
    }
    arr[high], arr[low] = arr[low], arr[high]
    return low
}

func quickSort(arr []int, low, high int) {
    if low < high {
        p := partition(arr, low, high)
        quickSort(arr, low, p-1)
        quickSort(arr, p+1, high)    
    }
}

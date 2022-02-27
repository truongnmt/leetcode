// Input: [1, 2, 1, 2, 2, 3, 3, 2, 3]
//         l              r
// Output: 6
// Max slice is [2, 2, 3, 3, 2, 3] which contains only two numbers 2 and 3 and the length is 6 

// mapNum tracking num from left to right
// loop right=0, right<len(nums),right++
//  if len(mapNum) == 2 and current number not exist in mapNum
//   loop in mapNum delete number that not equal right and right -1
//   set left = deleted number index+1
//  else 
//   mapNum[num] = i
//  compare right-left with max

func main() {
    nums := []int{1, 2, 1, 2, 2, 3, 3, 2, 3}
    fmt.Println(longestBiValueSlide(nums))
}

func longestBiValueSlide(nums []int) []int {
    mapNum := make(map[int]int)
    left, right := 0, 0
    max := 0
    maxIndex := []int{}
    
    for right=0;right<len(nums);right++ {
        num := nums[right]
        _, mapNumExist := mapNum[num]
        if len(mapNum) == 2 && !mapNumExist {
            for num, index := range mapNum {
                if num != nums[right] && num != nums[right-1] {
                    delete(mapNum, num)
                    left = index+1
                }
            }
            mapNum[num] = right
        } else {
            mapNum[num] = right
        }
        
        if (right - left) > max { 
            max = right - left
            maxIndex = []int{left, right}
        }
        fmt.Printf("mapNum=%v left=%v right=%v max=%v maxIndex=%v\n", mapNum, left, right, max, maxIndex)
    }
    
    return nums[maxIndex[0]:maxIndex[1]+1]
}

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSumHashTable(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumHashTable(nums []int, target int) []int {
	myMap := make(map[int]int)
  for index, num := range nums {
    targetSubstractbyNum := target - num
    if result, booleanCheck := myMap[targetSubstractbyNum]; booleanCheck {
      return []int{result, index}
    }
    myMap[num] = index
  }
  return []int{}
}

func twoSumHashTableRefactor(nums []int, target int) []int {
  m := make(map[int]int)
  for i, n := range nums {
    idx, ok := m[target - n]
    if ok {
      return []int{i, idx}
    }
    m[n] = i
  }
  return []int{}
}

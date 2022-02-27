// map[a[i]] = mean
// or
// map[mean] = [][a[i]]

package main

import (
	"fmt"
	"sort"
)

func main() {
    num := [][]int{
        {3, 3, 2, 4},
        {4, 4},
        {1, 4},
        {2, 5, 3, 0},
        {10, 2},
	}

    fmt.Println(solution(num))
}

func solution(a [][]int) [][]int {
    mapMean := make(map[float32][]int)
    mapOrder := []float32{}
    
    for i, row := range a {
        sumRow := 0
        for _, val := range row {
            sumRow += val
        }
        // fmt.Println(sumRow)
        // fmt.Println(float32(sumRow)/float32(len(row)))
        mapMean[float32(sumRow)/float32(len(row))] = append(mapMean[float32(sumRow)/float32(len(row))], i)
        mapOrder = append(mapOrder, float32(sumRow)/float32(len(row)))
	}

	fmt.Println(mapMean)

	keys := make([]float64, 0)
	for k, _ := range mapMean {
		keys = append(keys, float64(k))
	}
	sort.Float64s(keys)

	ans := [][]int{}
	for _, k := range keys {
		ans = append(ans, mapMean[float32(k)])
	}
    
    return ans
}

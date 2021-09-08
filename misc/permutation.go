// Input: n, k
// n is weight from 1 kg to 3^n kg. [1 3 9 27 ...]
// k is a things that weight k kg
// how to set k and weight so that it balance
// Eg: 
// k = 4: [4]         ----- 1 + 3
// k = 5: 1 + 3 + [5] ----- 9
// k = 6: [6] + 3     ----- 9

package main

import "fmt"

func main() {
	n := 100
	a := make([]int, n)
	binaryPrint(0, n, a)
}


func choose(x, n int, a []int) {
	if x == n {
		fmt.Println(a)
		return
	}
	for i:=0;i<2;i++ {
		a[x] = i
		binaryPrint(x+1, n, a)
	}
}

func getPermutations(elements []int) [][]int {
    permutations := [][]int{}
    if len(elements) == 1 {
        permutations = [][]int{elements}
        return permutations
    }
    for i := range elements {
        el := make([]int, len(elements))
        copy(el, elements)
        for _, perm := range getPermutations(append(el[0:i], el[i+1:]...)) {
            permutations = append(permutations, append([]int{elements[i]}, perm...))
        }
    }
    return permutations
}

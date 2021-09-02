// Input: n = 3
// Print an binary array has n element
// [0, 0, 0]
// [0, 0, 1]
// [0, 1, 0]
// [0, 1, 1]
// [1, 0, 0]
// [1, 0, 1]
// [1, 1, 0]
// [1, 1, 1]

package main

import "fmt"

func main() {
	n := 100
	a := make([]int, n)
	binaryPrint(0, n, a)
}


func binaryPrint(x, n int, a []int) {
	if x == n {
		fmt.Println(a)
		return
	}
	for i:=0;i<2;i++ {
		a[x] = i
		binaryPrint(x+1, n, a)
	}
}

// bP(0, 3, a) {
// 	i = 0
// 	a[0] = i = 0
// 	bP(1, 3, a) {
// 		i = 0
// 		a[1]=i=0
// 		bP(2, 3, a) {
// 			i = 0
// 			a[2] = i = 0
// 			bP(3, 3, a) {
// 				print a // [0 0 0]
// 			}

// 			i = 1
// 			a[2] = 1
// 			bP(3, 3, a) {
// 				print a // [0 0 1]
// 			}

// 			i = 2
// 			return
// 		}

// 		i = 1
// 		a[1] = i = 1
// 		bP(2, 3, a {
// 			a[2] = 0
// 			print a // [0 1 0]

// 			a[2] = 1
// 			print a // [0 1 1]
// 		})

// 		i=2 return
// 	}

// 	i=1
// 	a[0] = i = 1
// 	bP(1, 3, a) {
// 		i=0
// 		a[1] = 0
// 		bP(2, 3, a) {
// 			bP(3, 3, a) // [1 0 0]
// 			bP(3, 3, a) // [1 0 1]
// 		}

// 		i=1
// 		a[2] = 1
// 		bP(2, 3, a) {
// 			bP(3, 3, a) // [1 0 1]
// 			bP(3, 3, a) // [1 1 1]
// 		}
// 	}

// 	i=2
// 	return
// }

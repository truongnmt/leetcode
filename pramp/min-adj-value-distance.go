// Integer V lies strictly between integers U and W if U < V < W or if U > V > W.
// A non-empty zero-indexed array A consisting of N integers is given. A pair of 
// indices (P, Q), where 0 ≤ P < Q < N, is said to have adjacent values if no value 
// in the array lies strictly between values A[P] and A[Q].

// For example, in array A such that: A[0] = 0, A[1] = 3, A[2] = 3, A[3] = 7, A[4] = 5, A[5] = 3, A[6] = 11, A[7] = 1
// the following pairs of indices have adjacent values: 
// (0, 7), (1, 2), (1, 4), (1, 5), (1, 7), (2, 4), (2, 5), (2, 7), (3, 4), (3, 6), (4, 5), (5, 7)

// For example, indices 4 and 5 have adjacent values because there is no value 
// in array A that lies strictly between A[4] = 5and A[5] = 3; the only such value could be the number 4,
// and it is not present in the array.
// Given two indices their distance is defined as abs(P-Q). Create a function that 
// return the minimum distance between indices in an array that have adjacent values, if not return -1.

// Example:
// A = [1,4,7,3,3,5]
// Answere: 2 (p = 3 and q = 1)

// A = [1,4,7,3,3,5]
//     [1,3,3,4,5,7]
//      l   r
// idx [0,3,4,1,5,2]

// [0,3,3,7,5,3,11,1]
// [0,1,3,3,3,  5,7,11]

// [0,7,[1,2,5],4,3,6]

// [1, 2, 3, 4, 5]

// loop in a, luu map a[i]:[index min, index max]
// sort a -> sortedA
// loop sortedA
// shortestDistance = -1
// neu sortA != sortA+1
//   min(abs(map[sortA] - map[sortA]+1), shortestDistance) => sot truong hop

// A = [1,4,7,3,3,3,3,3,5]
// sort[1,3,3,4,5,7]
// idx [0,[3,4],[3,4],1,5,2]

func main() {
	// A := []int{1,4,7,3,3,5}
	// A := []int{0,3,3,7,5,3,11,1}
	A := []int{0, 3, 3, 3, 3, 3, 3, 3}
	fmt.Println(minAdjValueDistance(A))
}

func minAdjValueDistance(nums []int) int {
	m := make(map[int][2]int)
	for i, num := range nums {
		if _, exist := m[num]; exist {
			m[num] = [2]int{m[num][0], i}
		} else {
			m[num] = [2]int{i}
		}
	}

	sortedNums := nums
	sort.Ints(sortedNums)

	fmt.Println(sortedNums)
	fmt.Println(m)

	shortestDist := -1
	for i := 0; i < len(sortedNums); i++ {
		if i+1 < len(sortedNums) && sortedNums[i] != sortedNums[i+1] {
			indexI := m[sortedNums[i]]
			indexI1 := m[sortedNums[i+1]]

			var dist int
			if indexI[1] == 0 && indexI1[1] == 0 {        // [1,0] [2,0]
				dist = abs(indexI[0] - indexI1[0])
			} else if indexI[1] == 0 && indexI1[1] != 0 { // [1,0] [3,4]
				dist = min(abs(indexI[0]-indexI1[1]), abs(indexI[0]-indexI1[0]))
			} else if indexI[1] != 0 && indexI1[1] == 0 { // [3,4] [1,0]
				dist = min(abs(indexI[0]-indexI1[0]), abs(indexI[1]-indexI1[0]))
			} else {                                      // [5,6] [3,4]
				dist = min(abs(indexI[0]-indexI1[1]), abs(indexI[1]-indexI1[0]))
			}
			fmt.Printf("indexI=%v indexI1=%v dist=%v\n", indexI, indexI1, dist)

			if shortestDist == -1 && dist != 0 {
				shortestDist = dist
				continue
			}
			shortestDist = min(shortestDist, dist)
		}
	}

	return shortestDist
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

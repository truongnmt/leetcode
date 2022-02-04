// A 00 20
// B 01
// C 02 12
// D 21
// E 22 03 23

// put word to map

// directions up right down left
// loop in row
// loop in column

// func dfs(x, y, index int)
// if index == len(word) return true
// if neighbors is char
// return dfs (neighbor, index++)

// dfs(x,y, 0)

func exist(board [][]byte, word string) bool {
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	seen := make(map[[2]int]bool)

	var dfs func(int, int, int) bool
	dfs = func(x, y, index int) bool {
		// fmt.Printf("\nx=%v y=%v index=%v char=%v seen=%v\n", x, y, index, string(word[index]), seen)
		seen[[2]int{x, y}] = true
		if index == len(word) {
			return true
		}

		for _, drc := range directions {
			neighborX := x + drc[0]
			neighborY := y + drc[1]
			isSeen, _ := seen[[2]int{neighborX, neighborY}]
			// fmt.Printf("neighborX=%v neighborY=%v isSeen=%v\n", neighborX, neighborY, isSeen)
			if 0 <= neighborX && neighborX < len(board) && 0 <= neighborY && neighborY < len(board[0]) &&
				!isSeen && board[neighborX][neighborY] == word[index] {
				if dfs(neighborX, neighborY, index+1) == true {
					return true
				}
			}
		}
		seen[[2]int{x, y}] = false
		return false
	}

	for x, rows := range board {
		for y, val := range rows {
			if val == word[0] {
				if dfs(x, y, 1) == true {
					return true
				}
			}
			seen = make(map[[2]int]bool)
		}
	}

	return false
}

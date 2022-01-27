func setZeroes(matrix [][]int) {
	//     Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	//                    col
	//                       row
	//     Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

	//     loop to get 0 -> get 0 cols and 0 rows
	//     0 cols: 0 3
	//     0 rows: 0

	//     loop in line
	//      loop in row
	//      if line in 0 lines or row in 0 rows, matrix[line][row] = 0

	zeroCol := make(map[int]bool)
	zeroRow := make(map[int]bool)
	for rowNum, row := range matrix {
		for colNum, val := range row {
			if val == 0 {
				zeroCol[colNum] = true
				zeroRow[rowNum] = true
			}
		}
	}

	for rowNum, row := range matrix {
		for colNum, _ := range row {
			detectedZeroCol, _ := zeroCol[colNum]
			detectedZeroRow, _ := zeroRow[rowNum]
			if detectedZeroCol || detectedZeroRow {
				matrix[rowNum][colNum] = 0
			}
		}
	}
}

// if num positive
// divide to get the arrNum
// get the second smalletstNumber
// sortArrNum without smallestNumber
// else divide get the arrNum, sort decending

func smallestNumber(num int64) int64 {
	isMinus := false
	if num < 0 {
		isMinus = true
	}

	numStr := strconv.Itoa(int(math.Abs(float64(num))))
	numStrs := strings.Split(numStr, "")
	sort.Slice(numStrs, func(i, j int) bool {
		if isMinus {
			return numStrs[i] > numStrs[j]
		} else {
			return numStrs[i] < numStrs[j]
		}
	})

	smallestNonZeroIdx := 0
	for i, numStr := range numStrs {
		if numStr != "0" {
			smallestNonZeroIdx = i
			break
		}
	}
	numStrs[0], numStrs[smallestNonZeroIdx] = numStrs[smallestNonZeroIdx], numStrs[0]

	result, _ := strconv.Atoi(strings.Join(numStrs, ""))
	if isMinus {
		result *= -1
	}

	return int64(result)
}

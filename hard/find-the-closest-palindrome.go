// 123  121
// 1    0
// 30   33
// 26   22
// 20   22
// 16   11
// 11   9
// 10   9
// 9    8
// 1234 1221
// 999  1001
// 1000 999

// 42 44 33
// 100 99 101
// 253 252
// 259 262

// if 11 10 return 9
// if <10 return n-1

// find palin for 12345
// 12321
// 12421
// 12221
// 9999
// 10001
// get min of all palin

func nearestPalindromic(n string) string {

	// parse left part
	len := len(n)
	leftPart := 0
	if len%2 == 0 {
		leftPart, _ = strconv.Atoi(n[0 : len/2])
	} else {
		leftPart, _ = strconv.Atoi(n[0 : len/2+1])
	}
	nInt, _ := strconv.Atoi(n)

	// fmt.Printf("leftPart=%v nInt=%v\n", leftPart, nInt)

	// generate 5 palin
	palins := []int{}
	palins = append(palins, generatePalin(leftPart, len%2 == 0))   // 12321
	palins = append(palins, generatePalin(leftPart+1, len%2 == 0)) // 12421
	palins = append(palins, generatePalin(leftPart-1, len%2 == 0)) // 12221
	palins = append(palins, int(math.Pow(10, float64(len-1))-1))   // 9999
	palins = append(palins, int(math.Pow(10, float64(len))+1))     //   10001

	// compare diff with n
	minDiff := -1
	var ans int
	for _, palin := range palins {
		if palin == nInt {
			continue
		}
		diff := int(math.Abs(float64(palin - nInt)))
		if minDiff == -1 || diff < minDiff {
			minDiff = diff
			ans = palin
		} else if diff == minDiff {
			ans = int(math.Min(float64(ans), float64(palin)))
		}
	}

	return strconv.Itoa(ans)
}

func generatePalin(leftPart int, even bool) int {
	palin := leftPart
	if !even {
		leftPart = leftPart / 10
	}
	for leftPart > 0 {
		palin = palin*10 + leftPart%10
		leftPart /= 10
	}
	return palin
}

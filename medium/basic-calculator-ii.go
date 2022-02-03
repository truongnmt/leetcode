// 3+2*2

// 5+3*2-6/2-1+1
// 532*+62/1-1+

// Naive
// loop in s to split to number and operator, remove space
// loop again to calculate * and / first
// loop again to calculate + and -

func calculate(s string) int {
	// 3+2*4+4-2
	var splitedS []string
	number := ""
	for _, char := range s {
		if char == ' ' {
			continue
		}
		if 0 <= char-'0' && char-'0' <= '9' {
			number += string(char)
		} else {
			splitedS = append(splitedS, number)
			splitedS = append(splitedS, string(char))
			number = ""
		}
	}
	if number != "" {
		splitedS = append(splitedS, number)
	}
	// fmt.Println(splitedS)

	if len(splitedS) == 1 {
		tmp, _ := strconv.Atoi(splitedS[0])
		return tmp
	}

	// var splitedS2 []string
	for i := 0; i < len(splitedS); i++ {
		// fmt.Printf("i=%v splitedS[i]=%v\n", i, splitedS[i])
		if splitedS[i] == "*" || splitedS[i] == "/" {
			left, _ := strconv.Atoi(string(splitedS[i-1]))
			right, _ := strconv.Atoi(string(splitedS[i+1]))
			var result int
			if splitedS[i] == "*" {
				result = left * right
			}
			if splitedS[i] == "/" {
				result = left / right
			}
			splitedS[i-1] = strconv.Itoa(result)
			// fmt.Printf("result=%v splitedS=%v\n", result, splitedS)

			splitedS = append(splitedS[:i], splitedS[i+2:]...)
			// fmt.Println(splitedS)
			i -= 1
		}
	}

	result, _ := strconv.Atoi(splitedS[0])
	for i := 1; i < len(splitedS); i += 2 {
		if splitedS[i] == "+" {
			tmp, _ := strconv.Atoi(splitedS[i+1])
			result += tmp
		} else if splitedS[i] == "-" {
			tmp, _ := strconv.Atoi(splitedS[i+1])
			result -= tmp
		}
	}

	return result
}

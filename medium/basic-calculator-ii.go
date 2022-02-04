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

// ------------------------------------------------------------------------------------
// 5+3*2*4-6/2-1+1
// 532*4*+62/-1-1+
// 
// Reverse polish notation
// arrS []string
// loop in s, if number append to arrS, if operator put to stack
//   if priority operator > head, append operator to stack
//   else if <= priority, pop stack and append to arrS while operator bigger than head, then put operator to stack
// pop stack and append to arrS

// 3+2*4/5+3-1
// 324*5/+3+1-
// -
// 5 3 2*+6 2/-1-1+
// loop s in arrS
//   if s is number, convert to number, put to stack
//   if s is operator, pop 2 time in stack, calculate it and push result back to stack

// pop stack to get result

func isBiggerOperator(op1, op2 string) bool {
    op1Num := 0
    op2Num := 0
    if op1 == "*" || op1 == "/" { op1Num = 1}
    if op2 == "*" || op2 == "/" { op2Num = 1}
    if op1 == "+" || op1 == "-" { op1Num = -1}
    if op2 == "+" || op2 == "-" { op2Num = -1}
    
    return op1Num > op2Num
}

func calculate(s string) int {
    var arrS []string
    var stack []string
    number := ""
    for _, char := range s {
        if char == ' ' { continue }
        if 0 <= char - '0' && char - '0' <= '9' {
            number += string(char)
        } else {
            arrS = append(arrS, number)
            number = ""
            
            if len(stack) > 0 && isBiggerOperator(string(char), stack[len(stack)-1]) {
                stack = append(stack, string(char))
            } else {
                for len(stack) > 0 && !isBiggerOperator(string(char), stack[len(stack)-1]) {
                    // pop and push to arrS
                    arrS = append(arrS, stack[len(stack)-1])
                    stack = stack[:(len(stack)-1)]
                }
                // put operator to stack
                stack = append(stack, string(char))
            }
            // fmt.Printf("arrS= %v stack=%v\n", arrS, stack)
        }
    }
    if number != "" { arrS = append(arrS, number) }
    for len(stack) > 0 {
        // pop and push to arrS
        arrS = append(arrS, stack[len(stack)-1])
        stack = stack[:(len(stack)-1)]
    }
    
    // fmt.Printf("arrS=%v\n", arrS)
    
    // 532*4*+62/-1-1+
    // loop s in arrS
    //   if s is number, convert to number, put to stack
    //   if s is operator, pop 2 time in stack, calculate it and push result back to stack
    
    stackResult := []int{}
    for _, s := range arrS {
        if s != "+" && s != "-" && s != "*" && s != "/" {
            tmp, _ := strconv.Atoi(s)
            stackResult = append(stackResult, tmp)
        } else {
            var result int
            right := stackResult[len(stackResult)-1]
            left := stackResult[len(stackResult)-2]
            stackResult = stackResult[:(len(stackResult)-2)]
            if s == "+" { result += left + right }
            if s == "-" { result += left - right }
            if s == "*" { result += left * right }
            if s == "/" { result += left / right }
            stackResult = append(stackResult, result)
        }
    }
    
    return stackResult[0]
}
// ------------------------------------------------------------------------------------
// Stack
// 5+3*2*4-6/2-1+1
// i

// stack: 5 24 -3 -1 +1
// curNumber: 
// operator:

// for ch in s 
// if ch is number, add to curNumber
// else if ch is operator
//  if operator is '+' push curNumber to stack
//  if operator is '-' push -curNumber to stack
//  if operator is '*' pop stack , calculate and push result to stack
//  operator = ch
//  curNumber = 0

func calculate(s string) int {
    stack := []int{}
    curNumber := 0
    operator := '+'
    
    for i, ch := range s {
        if isNumber(ch) {
            curNumber = curNumber*10 + int(ch-'0')
            // fmt.Println(curNumber)
        }
        if !isNumber(ch) && ch != ' ' || i == len(s)-1 {
            if operator == '+' { stack = append(stack, curNumber) }
            if operator == '-' { stack = append(stack, -curNumber) }
            if operator == '*' {
                left := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack = append(stack, left * curNumber) 
            }
            if operator == '/' {
                left := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack = append(stack, left / curNumber) 
            }
            operator = ch
            curNumber = 0
        } 
    }
    
    // fmt.Println(stack)
    
    result := 0
    for _, num := range stack {
        result += num
    }
    
    return result
}

func isNumber(ch rune) bool { return '0' <= ch && ch <= '9' }

// ------------------------------------------------------------------------------------
// Stack TC O(n) SC O(1)
// 5+3*2/3-1
//         i

// result    : 5 -> 5+2=7
// lastNumber: 2 -> 2-1=1
// curNumber : 1
// operator  : -

// for ch in s 
// if ch is number, add to curNumber
// else if ch is operator
//  if operator is '+' || '-' result += lastNumber; lastnumber = operator curnumber
//  if operator is '*', lastnumber = lastnumber * curnumber
//  if operator is '/', lastnumber = lastnumber / curnumber
//  operator = ch
//  curNumber = 0

func calculate(s string) int {
    result := 0
    lastNumber := 0
    curNumber := 0
    operator := '+'
    
    for i, ch := range s {
        if isNumber(ch) {
            curNumber = curNumber*10 + int(ch-'0')
        }
        if !isNumber(ch) && ch != ' ' || i == len(s)-1 {
            if operator == '+' || operator == '-' { 
                result += lastNumber
                if operator == '+' { lastNumber = curNumber }
                if operator == '-' { lastNumber = -curNumber }
            }
            if operator == '*' {
                lastNumber = lastNumber * curNumber
            }
            if operator == '/' {
                lastNumber = lastNumber / curNumber
            }
            operator = ch
            curNumber = 0
        } 
    }
    
    result += lastNumber
    return result
}

func isNumber(ch rune) bool { return '0' <= ch && ch <= '9' }


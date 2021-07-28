package main

import "fmt"

func main() {
	s := "(())]"
	result := isValid(s)
	fmt.Println(result)
}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	m := make(map[string]string)
	m[")"] = "("
	m["}"] = "{"
	m["]"] = "["
	tmp := ""

	for _, ch := range s {
		if string(ch) == "(" || string(ch) == "{" || string(ch) == "[" {
			tmp += string(ch)
		} else {
			if tmp == "" || m[string(ch)] != string(tmp[len(tmp)-1]) {
				return false
			}
			tmp = tmp[:len(tmp)-1]
		}
	}

	if tmp == "" {
		return true
	}
	return false
}

func isValidRefactor(s string) bool {
	stack := []byte{}
    
    for i := 0; i < len(s); i++ {
		switch s[i] {
			case '(', '[', '{':
				stack = append(stack, s[i])
			case ')', ']', '}':
				if len(stack) == 0 {
					return false
				}
				lastChar := stack[len(stack) - 1]
            	stack = stack[:len(stack) - 1]
				if s[i] == ')' && lastChar != '(' {
					return false
				} else if s[i] == ']' && lastChar != '[' {
					return false
				} else if s[i] == '}' && lastChar != '{' {
					return false
				}
		}
	}
	if len(stack) > 0 {
        return false
    }
    
    return true
}















package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 10
	result := isPalindrome(x)
	fmt.Println(result)
}

func isPalindrome(x int) bool {
    s := strconv.Itoa(x)
    for i:=0; i<(len(s)-1/2); i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    
    return true
}

func isPalindromeRevertInt(x int) bool {
	if x < 0 {
		return false 
	}

    rev := 0
    y := x
    for y != 0 {
        remainder := y % 10
        y = y / 10
        rev = rev*10 + remainder
    }
    return rev == x
}

func isPalindromeRevertIntHalf(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false 
	}

    rev := 0
    for x > rev {
        remainder := x % 10
        x /= 10
        rev = rev*10 + remainder
    }
    return rev == x || rev/10 == x
}


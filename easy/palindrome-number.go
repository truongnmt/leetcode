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

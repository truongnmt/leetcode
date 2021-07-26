package main

import (
	"fmt"
)

func main() {
	s := "IV"
	result := romanToInt(s)
	fmt.Println(result)
}

func romanToInt(s string) int {
	romanMap := make(map[string]int)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000
	
	result := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 || romanMap[string(s[i])] >= romanMap[string(s[i+1])] {
			result += romanMap[string(s[i])]
		} else {
			result -= romanMap[string(s[i])]
		}
	}
	return result
}

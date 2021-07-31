package main

import (
	"fmt"
	"strings"
)

func main() {
	// arr := []string{"x"}
	// str := "A"
	
	// arr := []string{"x", "y", "z"}
	// str := "xyyzyzyx"
	
	// arr := []string{"A"}
	// str := "A" // result "A"

	// arr := []string{"x","y","z","r"}
	// str := "xyyzyzyx" // result ""

	arr := []string{"A","B","C"}
	str := "ADOBECODEBANCDDD" // result "BANC"


	// arr := []string{"A","B","C","E","K","I"}
	// str := "KADOBECODEBANCDDDEI" // result "KADOBECODEBANCDDDEI"

	
	result := GetShortestUniqueSubstring2(arr, str)
	fmt.Println(result)
}

func GetShortestUniqueSubstring(arr []string, str string) string {
	if len(arr) > len(str) {
		return ""
	}

	sumArr := 0
	for i:=0;i<len(arr);i++ {
		sumArr += int(arr[i][0])
	}
	
	tmp := 0
	for i:=0;i<len(arr);i++ {
		tmp += int(str[i])
	}
	if tmp == sumArr {
		return str[0:len(arr)]
	}

    for i:=1;i<=len(str)-len(arr);i++ {
		tmp -= int(str[i-1])
		tmp += int(str[i+len(arr)-1])
		if tmp == sumArr {
			return str[i:i+len(arr)]
		}
	}
	return ""
}

func GetShortestUniqueSubstring2(arr []string, str string) string {
	if len(arr) > len(str) {
		return ""
	}

	for i:=len(arr)-1;i<len(str);i++ {
		for j:=0;j<len(str)-i;j++ {
			containsArrs := false
			// check if substring contain all arr
			for _, arrS := range arr {
				if strings.Contains(str[j:j+i+1], arrS) {
					containsArrs = true
				} else {
					containsArrs = false
					break
				}
			}
		
			if containsArrs == true {
				return str[j:j+i+1]
			}			
		}
	}

	return ""
}

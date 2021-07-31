package main

import "fmt"

func main() {
	// arr := []string{"x"}
	// str := "A"
	// arr := []string{"x", "y", "z"}
	// str := "xyyzyzyx"
	arr := []string{"A","B","C"}
	str := "ADOBECODEBANCDDD"
	result := GetShortestUniqueSubstring(arr, str)
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

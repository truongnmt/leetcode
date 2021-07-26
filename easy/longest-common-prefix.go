package main

import "fmt"

func main() {
	strs := []string{"flower","flow","flight"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix 
	}

	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || ch != strs[j][i] {
				return prefix
			}
		}
		prefix += string(ch)
	}

	return prefix
}

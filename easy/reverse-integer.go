package main

import (
	"fmt"
	"strconv"
  "math"
)

func main() {
	x := -123
	result := reverse(x)
	fmt.Println(result)
}

func reverse(x int) int {
	if x == 0 {
		return x
	}

	s := strconv.Itoa(x)
	b := make([]byte, len(s))
	j := len(s) - 1
  inc := 0
	if x < 0 {
		inc = 1
    b[0] = s[0]
	}

	for i := 0+inc; i <= j; i++ {
		b[j-i+inc] = s[i]
	}

  result, _ := strconv.Atoi(string(b))
  if float64(result) > (math.Pow(2, 31)-1) || float64(result) < (math.Pow(-2, 31)) {
    return 0
  }

	return result
}

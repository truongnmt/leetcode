package main

import "fmt"

func main() {
	position := 9
	result := fiboAtPosition(position)
	fmt.Println(result)
}

func fiboAtPosition(p int) int {
	if p == 0 {
		return 0
	}
	if p == 1 {
		return 1
	}

	return fiboAtPosition(p - 2) + fiboAtPosition(p - 1)
}
